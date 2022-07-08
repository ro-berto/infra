// Code generated by cproto. DO NOT EDIT.

package api

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"ctrv2.api.CrosToolRunnerContainerService",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 164, 88, 95, 115, 219, 198,
			17, 231, 221, 129, 20, 120, 250, 203, 141, 164, 40, 176, 37, 173,
			105, 197, 97, 82, 17, 210, 48, 19, 39, 145, 39, 238, 88, 138,
			210, 40, 246, 88, 13, 101, 53, 110, 213, 218, 129, 136, 35, 137,
			49, 9, 48, 184, 131, 92, 125, 151, 62, 245, 165, 47, 237, 67,
			190, 65, 167, 31, 160, 211, 247, 126, 153, 206, 29, 112, 36, 101,
			203, 174, 167, 125, 227, 30, 22, 191, 219, 253, 221, 111, 119, 15,
			228, 63, 111, 241, 123, 81, 220, 77, 131, 157, 78, 154, 200, 157,
			206, 48, 52, 63, 154, 42, 73, 6, 205, 52, 139, 99, 145, 238,
			4, 163, 104, 167, 147, 196, 42, 136, 98, 145, 62, 151, 34, 189,
			136, 58, 194, 31, 165, 137, 74, 160, 218, 81, 233, 69, 203, 15,
			70, 81, 253, 128, 207, 60, 22, 234, 101, 146, 190, 0, 224, 78,
			28, 12, 197, 26, 65, 210, 168, 182, 205, 111, 88, 224, 52, 10,
			215, 168, 89, 161, 81, 8, 203, 188, 156, 188, 140, 69, 184, 198,
			144, 52, 220, 118, 110, 212, 15, 121, 245, 192, 110, 245, 127, 192,
			124, 194, 151, 15, 82, 17, 40, 81, 68, 212, 22, 63, 101, 66,
			170, 235, 16, 235, 135, 124, 229, 21, 95, 57, 74, 98, 41, 96,
			155, 207, 196, 249, 146, 241, 159, 109, 129, 63, 206, 214, 183, 206,
			214, 165, 254, 17, 175, 253, 74, 168, 119, 216, 111, 159, 195, 180,
			227, 255, 180, 89, 141, 47, 158, 244, 51, 21, 38, 47, 227, 98,
			171, 58, 240, 165, 201, 82, 14, 90, 255, 11, 229, 43, 39, 42,
			72, 213, 152, 211, 183, 4, 6, 31, 241, 197, 201, 49, 71, 195,
			160, 39, 10, 158, 23, 198, 203, 71, 122, 21, 126, 224, 16, 132,
			97, 164, 162, 36, 14, 6, 207, 147, 145, 254, 33, 205, 1, 204,
			182, 26, 83, 97, 95, 187, 181, 127, 156, 251, 183, 107, 19, 140,
			98, 9, 110, 243, 121, 169, 223, 121, 222, 73, 134, 195, 32, 14,
			215, 28, 100, 141, 106, 123, 78, 230, 64, 102, 205, 59, 225, 51,
			214, 127, 149, 87, 196, 31, 71, 137, 212, 121, 104, 199, 194, 210,
			235, 23, 201, 32, 27, 234, 4, 204, 122, 110, 193, 218, 132, 100,
			102, 50, 27, 19, 250, 136, 175, 190, 26, 109, 113, 48, 45, 94,
			29, 167, 95, 28, 205, 242, 84, 142, 147, 23, 38, 110, 173, 127,
			83, 190, 113, 144, 38, 242, 73, 146, 12, 218, 166, 136, 198, 110,
			39, 121, 249, 64, 155, 207, 95, 81, 29, 108, 78, 131, 94, 163,
			93, 15, 223, 236, 80, 132, 122, 196, 249, 68, 89, 112, 115, 202,
			255, 53, 101, 122, 235, 111, 120, 90, 64, 29, 112, 215, 170, 9,
			188, 233, 35, 189, 170, 58, 239, 198, 181, 207, 10, 144, 83, 190,
			112, 149, 84, 192, 255, 166, 14, 239, 214, 91, 60, 114, 216, 253,
			15, 127, 119, 251, 29, 90, 214, 119, 255, 244, 120, 5, 28, 167,
			212, 37, 252, 111, 132, 147, 57, 96, 78, 9, 90, 127, 38, 120,
			144, 140, 46, 211, 168, 215, 87, 216, 218, 109, 181, 240, 73, 95,
			224, 65, 63, 77, 134, 81, 54, 196, 227, 19, 124, 144, 169, 126,
			146, 74, 31, 31, 12, 6, 104, 252, 36, 166, 66, 183, 61, 17,
			250, 28, 79, 165, 192, 164, 139, 170, 31, 73, 148, 73, 150, 118,
			4, 118, 146, 80, 96, 36, 177, 151, 92, 136, 52, 22, 33, 158,
			95, 98, 128, 251, 39, 95, 55, 165, 186, 28, 8, 28, 68, 29,
			17, 75, 129, 170, 31, 40, 236, 4, 49, 158, 11, 142, 221, 36,
			139, 67, 140, 98, 84, 125, 129, 143, 142, 14, 14, 31, 159, 28,
			98, 55, 26, 8, 159, 115, 151, 19, 10, 172, 82, 2, 253, 203,
			5, 230, 150, 246, 120, 149, 83, 119, 54, 255, 249, 5, 167, 149,
			18, 56, 179, 165, 77, 226, 109, 227, 161, 209, 188, 68, 93, 51,
			73, 140, 97, 210, 121, 33, 82, 44, 42, 72, 98, 32, 177, 104,
			217, 146, 115, 206, 89, 165, 68, 128, 205, 186, 119, 120, 131, 59,
			149, 18, 45, 129, 51, 79, 23, 182, 60, 15, 115, 77, 73, 12,
			44, 68, 81, 27, 156, 207, 241, 178, 246, 36, 192, 230, 43, 239,
			89, 139, 2, 155, 95, 110, 88, 139, 1, 91, 152, 175, 243, 150,
			193, 36, 224, 212, 40, 220, 242, 182, 176, 45, 84, 26, 137, 11,
			33, 49, 138, 187, 137, 230, 237, 77, 232, 132, 0, 171, 85, 150,
			172, 69, 129, 213, 106, 117, 107, 49, 96, 48, 191, 201, 119, 13,
			58, 5, 103, 133, 174, 110, 122, 117, 212, 130, 147, 168, 37, 135,
			7, 79, 218, 56, 174, 63, 155, 176, 197, 166, 4, 216, 74, 101,
			193, 90, 20, 216, 202, 226, 134, 181, 24, 176, 213, 249, 117, 254,
			119, 106, 192, 25, 56, 235, 116, 227, 67, 239, 103, 138, 237, 44,
			158, 34, 99, 2, 254, 50, 82, 125, 115, 104, 163, 52, 185, 136,
			66, 17, 162, 105, 79, 150, 113, 159, 227, 19, 173, 141, 64, 202,
			108, 40, 164, 125, 95, 175, 12, 82, 17, 132, 151, 24, 100, 170,
			47, 98, 21, 117, 2, 37, 66, 84, 9, 142, 178, 193, 192, 32,
			202, 108, 52, 26, 68, 34, 68, 211, 125, 13, 146, 184, 178, 179,
			86, 100, 22, 107, 213, 132, 66, 5, 157, 190, 8, 113, 168, 197,
			215, 104, 134, 31, 223, 195, 96, 48, 192, 188, 3, 134, 56, 74,
			82, 37, 243, 55, 180, 218, 70, 217, 249, 32, 146, 253, 124, 191,
			0, 211, 32, 14, 147, 161, 113, 194, 36, 198, 126, 34, 21, 54,
			154, 191, 214, 24, 113, 104, 66, 121, 101, 215, 115, 129, 169, 24,
			38, 23, 34, 228, 24, 116, 149, 78, 72, 161, 84, 201, 72, 98,
			163, 217, 76, 135, 31, 251, 150, 108, 70, 128, 173, 87, 150, 173,
			69, 129, 173, 175, 124, 98, 45, 6, 108, 99, 254, 54, 255, 156,
			83, 167, 4, 78, 189, 244, 49, 241, 126, 129, 109, 49, 210, 245,
			21, 43, 137, 231, 129, 140, 58, 111, 209, 10, 231, 204, 209, 58,
			172, 187, 139, 252, 6, 119, 28, 173, 95, 182, 69, 193, 91, 192,
			162, 131, 161, 30, 101, 38, 22, 253, 176, 172, 159, 186, 214, 34,
			192, 182, 170, 243, 214, 98, 192, 182, 150, 106, 124, 219, 192, 16,
			96, 119, 232, 146, 183, 57, 134, 57, 250, 90, 159, 96, 212, 43,
			202, 57, 15, 196, 226, 146, 178, 118, 183, 184, 90, 186, 119, 170,
			179, 214, 98, 192, 238, 44, 44, 242, 83, 131, 75, 129, 53, 104,
			205, 251, 22, 143, 226, 208, 28, 184, 196, 168, 107, 8, 46, 114,
			194, 151, 129, 196, 142, 41, 61, 179, 145, 225, 62, 75, 83, 17,
			171, 215, 5, 237, 219, 0, 104, 89, 227, 86, 172, 69, 128, 53,
			102, 230, 172, 197, 128, 53, 22, 151, 248, 151, 156, 58, 4, 156,
			237, 210, 167, 196, 107, 190, 3, 201, 227, 221, 114, 154, 117, 86,
			219, 110, 141, 175, 115, 199, 33, 154, 102, 159, 130, 183, 132, 227,
			118, 60, 69, 52, 49, 68, 251, 5, 33, 196, 16, 237, 23, 68,
			19, 67, 180, 191, 84, 227, 190, 1, 34, 192, 118, 233, 146, 119,
			107, 10, 232, 45, 84, 19, 67, 245, 238, 24, 89, 7, 181, 91,
			80, 77, 12, 213, 187, 11, 139, 252, 7, 131, 76, 129, 181, 104,
			205, 251, 238, 117, 170, 167, 180, 172, 187, 161, 174, 214, 119, 39,
			155, 24, 178, 91, 5, 217, 196, 144, 221, 42, 200, 38, 134, 236,
			214, 226, 18, 231, 156, 58, 20, 156, 207, 74, 159, 19, 195, 158,
			246, 250, 204, 189, 201, 103, 185, 227, 80, 205, 222, 93, 10, 230,
			21, 106, 184, 186, 91, 100, 68, 13, 87, 119, 11, 174, 168, 225,
			234, 238, 82, 205, 192, 49, 112, 190, 44, 221, 203, 225, 116, 81,
			125, 233, 174, 27, 56, 166, 225, 246, 232, 170, 121, 133, 209, 82,
			69, 91, 85, 107, 17, 96, 123, 188, 102, 45, 6, 108, 111, 121,
			197, 192, 57, 224, 124, 85, 250, 101, 14, 231, 16, 96, 95, 185,
			31, 24, 56, 71, 195, 221, 47, 162, 115, 76, 116, 247, 139, 232,
			28, 3, 119, 191, 136, 206, 49, 112, 247, 139, 232, 202, 224, 60,
			40, 29, 228, 112, 101, 2, 236, 129, 235, 25, 184, 178, 134, 219,
			47, 162, 43, 155, 232, 246, 139, 232, 202, 6, 110, 191, 136, 174,
			108, 224, 246, 151, 87, 244, 100, 115, 42, 192, 14, 75, 158, 65,
			171, 16, 96, 135, 238, 251, 102, 121, 6, 216, 55, 165, 27, 102,
			121, 134, 0, 251, 198, 93, 51, 155, 187, 224, 124, 91, 58, 207,
			55, 119, 9, 176, 111, 221, 117, 222, 229, 142, 227, 234, 205, 31,
			82, 240, 126, 139, 167, 113, 244, 83, 38, 140, 72, 177, 23, 93,
			136, 88, 55, 190, 171, 98, 48, 227, 216, 118, 183, 76, 55, 205,
			65, 160, 59, 155, 74, 48, 45, 230, 22, 159, 114, 215, 5, 83,
			104, 194, 53, 52, 61, 44, 104, 114, 77, 94, 15, 11, 154, 92,
			147, 215, 195, 165, 26, 255, 23, 49, 33, 17, 96, 199, 116, 221,
			251, 7, 193, 71, 73, 39, 208, 23, 87, 93, 118, 166, 195, 79,
			223, 8, 48, 140, 82, 209, 81, 131, 75, 51, 14, 166, 75, 193,
			199, 199, 137, 42, 124, 141, 120, 167, 163, 122, 199, 49, 227, 115,
			20, 126, 207, 199, 76, 54, 11, 208, 209, 139, 158, 31, 138, 139,
			252, 238, 148, 138, 94, 36, 85, 122, 185, 163, 132, 84, 77, 123,
			95, 200, 159, 133, 153, 218, 211, 196, 72, 101, 147, 215, 53, 121,
			60, 78, 94, 215, 228, 113, 117, 205, 90, 12, 216, 241, 141, 155,
			252, 72, 231, 206, 74, 224, 124, 79, 127, 207, 188, 123, 120, 146,
			141, 244, 180, 17, 33, 22, 223, 11, 56, 12, 84, 167, 143, 157,
			36, 77, 205, 197, 46, 140, 226, 30, 254, 88, 164, 145, 102, 241,
			143, 216, 29, 4, 61, 105, 25, 103, 154, 227, 239, 121, 141, 255,
			149, 240, 138, 54, 245, 81, 159, 58, 155, 222, 159, 72, 113, 3,
			50, 211, 204, 71, 41, 4, 78, 80, 176, 217, 204, 39, 98, 126,
			157, 187, 118, 58, 78, 134, 35, 191, 118, 58, 106, 213, 232, 9,
			233, 227, 65, 222, 43, 6, 151, 152, 196, 131, 75, 51, 174, 115,
			31, 113, 5, 217, 146, 253, 197, 46, 231, 11, 124, 38, 143, 214,
			209, 225, 206, 77, 236, 50, 176, 211, 249, 247, 38, 54, 1, 118,
			186, 236, 77, 108, 6, 236, 116, 125, 131, 103, 69, 182, 4, 216,
			83, 103, 211, 235, 226, 111, 204, 167, 12, 14, 147, 44, 86, 81,
			220, 187, 38, 225, 252, 99, 167, 136, 97, 71, 13, 71, 59, 58,
			250, 166, 76, 59, 197, 69, 89, 72, 181, 103, 214, 199, 50, 106,
			134, 66, 170, 201, 211, 73, 216, 196, 209, 251, 78, 194, 214, 71,
			255, 116, 42, 108, 125, 248, 79, 167, 194, 214, 199, 255, 116, 125,
			131, 159, 21, 97, 83, 96, 103, 206, 251, 222, 35, 221, 238, 99,
			209, 81, 182, 10, 117, 93, 134, 175, 140, 247, 107, 82, 177, 131,
			63, 207, 229, 60, 141, 194, 158, 152, 4, 167, 27, 243, 153, 51,
			101, 19, 96, 103, 179, 48, 177, 25, 176, 179, 149, 85, 254, 169,
			169, 67, 10, 236, 25, 189, 229, 221, 193, 7, 227, 239, 206, 177,
			22, 187, 73, 122, 69, 123, 227, 58, 167, 21, 253, 86, 213, 90,
			4, 216, 51, 126, 211, 90, 12, 216, 179, 77, 228, 207, 13, 60,
			3, 22, 208, 45, 175, 141, 197, 119, 170, 105, 35, 89, 174, 30,
			243, 241, 160, 59, 8, 6, 147, 210, 181, 42, 57, 171, 219, 50,
			171, 111, 99, 189, 169, 5, 164, 127, 124, 177, 91, 223, 70, 223,
			247, 255, 96, 67, 97, 142, 222, 97, 108, 149, 129, 5, 179, 53,
			107, 17, 96, 1, 108, 90, 75, 199, 82, 191, 109, 90, 101, 21,
			156, 80, 127, 2, 233, 86, 89, 37, 192, 66, 119, 195, 244, 233,
			170, 174, 31, 65, 61, 243, 74, 213, 244, 105, 65, 103, 173, 69,
			128, 137, 185, 21, 107, 49, 96, 98, 237, 131, 243, 138, 249, 203,
			231, 211, 255, 4, 0, 0, 255, 255, 86, 214, 249, 45, 52, 18,
			0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptorpb.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("ctrv2.api.CrosToolRunnerContainerService")
	if err != nil {
		panic(err)
	}
	return ret
}
