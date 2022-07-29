package builtins

import (
	"context"
	"crypto"
	"encoding/binary"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"infra/libs/cipkg"
)

const CopyFilesBuilder = BuiltinBuilderPrefix + "copyFiles"

var (
	copyFilesHashMap       = make(map[string]fs.FS)
	copyFilesHashAlgorithm = crypto.SHA256
)

type CopyFiles struct {
	Name  string
	Files fs.FS
}

func (cf *CopyFiles) Generate(ctx *cipkg.BuildContext) (cipkg.Derivation, cipkg.PackageMetadata, error) {
	h := copyFilesHashAlgorithm.New()
	fs.WalkDir(cf.Files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Hash path
		if _, err := h.Write([]byte(path)); err != nil {
			return fmt.Errorf("write path failed: %s: %w", path, err)
		}

		if d.IsDir() {
			return nil
		}

		// Hash file content
		f, err := cf.Files.Open(path)
		if err != nil {
			return fmt.Errorf("open file failed: %s: %w", path, err)
		}
		if _, err := io.Copy(h, f); err != nil {
			return fmt.Errorf("write file failed: %s: %w", path, err)
		}

		// Hash file mode
		var mode [4]byte
		finfo, err := f.Stat()
		if err != nil {
			return fmt.Errorf("get file info failed: %s: %w", path, err)
		}
		binary.LittleEndian.PutUint32(mode[:], uint32(finfo.Mode()))
		if _, err := h.Write(mode[:]); err != nil {
			return fmt.Errorf("write mode failed: %s: %w", path, err)
		}

		return nil
	})
	hashValue := fmt.Sprintf("%s:%x", copyFilesHashAlgorithm, h.Sum(nil))
	copyFilesHashMap[hashValue] = cf.Files
	return cipkg.Derivation{
		Name:    cf.Name,
		Builder: CopyFilesBuilder,
		Args:    []string{hashValue},
	}, cipkg.PackageMetadata{}, nil
}

func copyFiles(ctx context.Context, cmd *exec.Cmd) error {
	// cmd.Args = ["builtin:copyFiles", filesHash]
	if len(cmd.Args) != 2 {
		return fmt.Errorf("invalid arguments: %v", cmd.Args)
	}
	out := GetEnv("out", cmd.Env)

	h := copyFilesHashAlgorithm.New()
	files := copyFilesHashMap[cmd.Args[1]]
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Hash path
		if _, err := h.Write([]byte(path)); err != nil {
			return fmt.Errorf("write path failed: %s: %w", path, err)
		}

		dst := filepath.Join(out, path)

		// Create path and return if it's directory
		if d.IsDir() {
			if err := os.MkdirAll(dst, os.ModePerm); err != nil {
				return fmt.Errorf("create dir failed: %s: %w", dst, err)
			}
			return nil
		}

		// Hash and copy file content
		dstFile, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("create dst file failed: %s: %w", dst, err)
		}
		srcFile, err := files.Open(path)
		if err != nil {
			return fmt.Errorf("open src file failed: %s: %w", dst, err)
		}

		if _, err := io.Copy(dstFile, io.TeeReader(srcFile, h)); err != nil {
			return fmt.Errorf("copy file failed: %s: %w", dst, err)
		}

		// Hash and copy file mode
		srcInfo, err := srcFile.Stat()
		if err != nil {
			return fmt.Errorf("get file info failed: %s: %w", path, err)
		}
		if err := dstFile.Chmod(srcInfo.Mode()); err != nil {
			return fmt.Errorf("chmod failed: %s: %w", dst, err)
		}

		var mode [4]byte
		binary.LittleEndian.PutUint32(mode[:], uint32(srcInfo.Mode()))
		if _, err := h.Write(mode[:]); err != nil {
			return fmt.Errorf("write mode failed: %s: %w", path, err)
		}

		return nil
	})

	hashValue := fmt.Sprintf("%s:%x", copyFilesHashAlgorithm, h.Sum(nil))
	if hashValue != cmd.Args[1] {
		return fmt.Errorf("hash value mismach: expected: %s, result: %s", cmd.Args[1], hashValue)
	}
	return nil
}