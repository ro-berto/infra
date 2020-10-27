// Code generated by cproto. DO NOT EDIT.

package api

import "go.chromium.org/luci/grpc/discovery"

import "google.golang.org/protobuf/types/descriptorpb"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"crrev.Crrev",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 204, 88, 77, 111, 28, 199,
			209, 222, 153, 158, 161, 151, 173, 47, 170, 45, 145, 203, 53, 205,
			183, 76, 189, 150, 72, 138, 220, 53, 41, 27, 132, 233, 128, 1,
			45, 89, 54, 3, 130, 32, 72, 73, 129, 3, 1, 194, 236, 108,
			239, 78, 155, 179, 211, 171, 238, 153, 165, 137, 192, 65, 124, 201,
			33, 57, 4, 72, 206, 65, 114, 74, 254, 128, 115, 10, 144, 67,
			126, 73, 126, 65, 110, 185, 7, 93, 211, 61, 92, 126, 56, 17,
			20, 8, 200, 109, 170, 167, 167, 251, 169, 170, 231, 169, 174, 30,
			250, 139, 85, 250, 161, 200, 122, 42, 106, 71, 195, 33, 207, 250,
			34, 227, 237, 88, 173, 42, 62, 106, 247, 148, 204, 114, 158, 117,
			219, 209, 80, 180, 71, 107, 109, 205, 213, 72, 196, 188, 53, 84,
			50, 151, 44, 140, 149, 226, 163, 133, 123, 244, 198, 1, 239, 10,
			197, 227, 252, 128, 191, 44, 184, 206, 217, 45, 26, 190, 44, 184,
			58, 105, 120, 224, 45, 78, 30, 148, 198, 194, 183, 30, 157, 58,
			157, 169, 135, 50, 211, 156, 205, 210, 122, 95, 228, 47, 146, 72,
			39, 118, 246, 91, 198, 214, 73, 196, 110, 211, 32, 145, 58, 111,
			248, 229, 240, 80, 201, 175, 120, 156, 179, 6, 165, 138, 15, 165,
			22, 185, 84, 39, 13, 130, 47, 3, 51, 194, 222, 163, 87, 149,
			93, 255, 69, 161, 210, 70, 128, 239, 174, 184, 177, 167, 42, 93,
			248, 141, 71, 167, 246, 138, 65, 135, 43, 145, 245, 29, 92, 102,
			55, 42, 247, 199, 103, 54, 127, 102, 151, 18, 194, 216, 136, 217,
			11, 159, 133, 204, 94, 40, 222, 179, 56, 174, 184, 177, 3, 222,
			99, 247, 232, 141, 106, 74, 134, 123, 34, 34, 114, 112, 221, 13,
			151, 72, 22, 126, 238, 209, 155, 99, 160, 254, 115, 100, 254, 239,
			226, 202, 62, 174, 60, 81, 90, 85, 232, 200, 191, 11, 93, 112,
			26, 186, 133, 101, 122, 237, 161, 28, 12, 68, 149, 194, 203, 119,
			255, 34, 210, 201, 194, 31, 60, 122, 221, 77, 126, 3, 89, 188,
			196, 185, 224, 140, 115, 231, 211, 28, 94, 72, 243, 250, 95, 60,
			26, 62, 52, 236, 100, 159, 208, 186, 227, 28, 155, 110, 33, 99,
			91, 231, 232, 218, 156, 185, 48, 110, 221, 218, 162, 147, 85, 94,
			152, 155, 117, 158, 62, 205, 198, 197, 23, 246, 251, 143, 232, 68,
			25, 40, 118, 203, 206, 57, 19, 228, 230, 237, 115, 163, 229, 103,
			159, 174, 252, 100, 249, 213, 4, 249, 73, 52, 20, 63, 250, 235,
			251, 116, 130, 5, 87, 106, 137, 71, 255, 73, 168, 119, 149, 145,
			43, 53, 182, 254, 123, 15, 30, 202, 225, 137, 18, 253, 36, 135,
			245, 15, 214, 63, 128, 39, 9, 135, 135, 137, 146, 3, 81, 12,
			96, 187, 200, 19, 169, 116, 11, 182, 211, 20, 112, 146, 6, 197,
			141, 190, 121, 183, 69, 225, 169, 230, 32, 123, 144, 39, 66, 131,
			150, 133, 138, 57, 196, 178, 203, 65, 104, 232, 203, 17, 87, 25,
			239, 66, 231, 4, 34, 248, 244, 240, 209, 170, 206, 79, 82, 14,
			169, 136, 121, 166, 57, 228, 73, 148, 67, 28, 101, 208, 225, 20,
			122, 178, 200, 186, 32, 50, 200, 19, 14, 187, 59, 15, 63, 219,
			59, 252, 12, 122, 34, 229, 45, 186, 254, 91, 31, 158, 152, 245,
			177, 152, 224, 32, 116, 121, 79, 100, 92, 67, 233, 45, 108, 239,
			239, 192, 104, 173, 5, 59, 185, 217, 184, 19, 105, 222, 5, 153,
			193, 80, 241, 145, 144, 133, 134, 17, 87, 90, 200, 12, 100, 143,
			194, 193, 103, 135, 79, 198, 191, 139, 204, 190, 57, 36, 145, 198,
			189, 251, 50, 74, 33, 151, 112, 196, 249, 16, 226, 84, 240, 44,
			215, 208, 43, 210, 244, 4, 98, 57, 24, 70, 185, 232, 24, 80,
			240, 149, 150, 25, 244, 4, 79, 187, 26, 34, 197, 161, 48, 123,
			230, 18, 180, 24, 12, 83, 209, 51, 179, 179, 211, 109, 33, 229,
			253, 40, 62, 129, 44, 26, 136, 172, 15, 58, 78, 248, 192, 172,
			130, 192, 215, 33, 227, 188, 171, 161, 195, 65, 100, 185, 146, 221,
			34, 46, 215, 234, 137, 175, 161, 151, 70, 199, 218, 4, 166, 244,
			17, 186, 92, 139, 126, 214, 162, 235, 15, 48, 35, 150, 32, 232,
			133, 163, 5, 12, 35, 21, 13, 120, 206, 85, 9, 77, 241, 151,
			133, 80, 38, 93, 180, 78, 61, 159, 145, 107, 181, 235, 230, 169,
			206, 200, 141, 218, 54, 157, 164, 126, 253, 74, 249, 72, 169, 31,
			212, 88, 112, 179, 54, 227, 81, 74, 73, 80, 243, 24, 185, 89,
			159, 161, 255, 240, 104, 16, 212, 252, 26, 35, 211, 254, 219, 205,
			191, 123, 128, 117, 218, 196, 154, 191, 44, 202, 120, 61, 61, 216,
			133, 97, 148, 39, 198, 89, 228, 42, 12, 34, 145, 129, 211, 26,
			164, 178, 47, 226, 22, 60, 150, 10, 248, 215, 209, 96, 152, 242,
			21, 16, 57, 181, 4, 128, 133, 246, 131, 5, 88, 172, 102, 231,
			18, 98, 203, 192, 182, 86, 177, 9, 253, 64, 228, 112, 44, 242,
			4, 156, 230, 225, 193, 18, 72, 69, 161, 253, 83, 45, 7, 252,
			69, 57, 5, 43, 203, 55, 103, 87, 50, 175, 225, 180, 134, 84,
			252, 160, 112, 254, 195, 165, 21, 224, 121, 220, 162, 244, 42, 13,
			141, 191, 161, 113, 184, 238, 44, 143, 145, 233, 201, 235, 206, 34,
			140, 76, 223, 100, 24, 50, 143, 5, 179, 181, 255, 47, 67, 230,
			121, 140, 204, 214, 27, 244, 5, 13, 2, 207, 68, 108, 206, 95,
			105, 30, 128, 43, 122, 38, 102, 17, 242, 201, 12, 57, 191, 240,
			141, 236, 193, 32, 202, 227, 132, 119, 221, 48, 42, 68, 232, 146,
			91, 61, 227, 171, 115, 203, 66, 244, 16, 226, 156, 133, 232, 33,
			196, 185, 201, 155, 206, 34, 140, 204, 221, 186, 237, 172, 58, 35,
			115, 211, 247, 157, 69, 25, 153, 155, 89, 30, 183, 222, 91, 166,
			31, 35, 104, 143, 145, 121, 127, 177, 185, 2, 166, 28, 151, 74,
			150, 253, 148, 91, 121, 227, 224, 34, 111, 245, 91, 85, 138, 150,
			42, 56, 94, 104, 190, 117, 112, 76, 44, 230, 39, 175, 57, 139,
			48, 50, 63, 229, 192, 25, 246, 205, 179, 123, 206, 162, 140, 204,
			191, 125, 119, 220, 122, 247, 46, 253, 33, 194, 241, 25, 1, 255,
			126, 115, 125, 60, 133, 66, 195, 231, 34, 31, 31, 57, 139, 200,
			144, 102, 201, 129, 242, 67, 179, 130, 3, 101, 220, 131, 201, 183,
			157, 69, 24, 129, 233, 25, 103, 213, 25, 129, 134, 139, 138, 79,
			25, 129, 217, 165, 113, 235, 206, 18, 253, 165, 135, 168, 8, 35,
			119, 252, 102, 243, 27, 24, 63, 97, 44, 46, 145, 114, 141, 130,
			192, 202, 200, 109, 58, 87, 240, 89, 71, 3, 142, 239, 48, 185,
			133, 230, 10, 132, 62, 77, 108, 41, 252, 227, 132, 103, 80, 104,
			83, 42, 80, 75, 247, 244, 165, 106, 114, 14, 146, 208, 128, 113,
			14, 18, 143, 145, 59, 147, 46, 237, 196, 0, 109, 204, 34, 79,
			125, 22, 220, 173, 173, 151, 60, 53, 97, 184, 91, 111, 96, 202,
			125, 195, 211, 69, 159, 189, 70, 202, 125, 100, 224, 162, 221, 220,
			71, 6, 46, 218, 148, 251, 200, 192, 197, 169, 155, 152, 72, 12,
			252, 178, 223, 120, 237, 68, 250, 200, 174, 229, 106, 43, 195, 174,
			101, 155, 72, 31, 217, 181, 60, 61, 67, 191, 243, 112, 47, 159,
			145, 150, 223, 108, 254, 209, 131, 241, 230, 203, 108, 151, 153, 4,
			200, 222, 105, 25, 41, 207, 18, 60, 124, 70, 81, 90, 224, 203,
			190, 200, 87, 123, 82, 230, 92, 225, 163, 30, 101, 171, 162, 75,
			65, 42, 120, 168, 86, 203, 67, 120, 117, 223, 45, 80, 34, 86,
			188, 167, 219, 9, 143, 186, 186, 61, 136, 116, 206, 213, 10, 5,
			61, 202, 54, 219, 109, 61, 202, 90, 206, 159, 150, 84, 253, 54,
			26, 188, 157, 171, 34, 59, 58, 227, 160, 97, 106, 171, 114, 208,
			4, 172, 101, 19, 233, 35, 83, 91, 141, 89, 218, 65, 255, 8,
			35, 107, 254, 92, 243, 41, 156, 235, 128, 140, 135, 218, 156, 9,
			89, 46, 162, 20, 68, 215, 60, 244, 4, 87, 88, 151, 203, 178,
			34, 50, 232, 139, 17, 207, 160, 163, 162, 44, 78, 40, 44, 142,
			199, 232, 52, 179, 134, 86, 107, 254, 91, 206, 242, 24, 89, 171,
			207, 56, 203, 0, 104, 190, 131, 180, 34, 44, 248, 176, 246, 105,
			73, 43, 51, 235, 195, 250, 44, 150, 63, 98, 104, 181, 241, 6,
			203, 31, 65, 242, 109, 216, 128, 17, 36, 223, 134, 45, 127, 4,
			201, 183, 97, 203, 31, 193, 242, 183, 97, 203, 31, 193, 130, 183,
			97, 203, 159, 181, 222, 91, 166, 191, 242, 16, 181, 199, 200, 166,
			191, 214, 252, 217, 133, 216, 150, 248, 52, 156, 111, 236, 206, 79,
			108, 193, 190, 146, 157, 168, 147, 158, 24, 236, 41, 215, 154, 226,
			137, 86, 118, 19, 43, 208, 41, 114, 56, 226, 195, 220, 56, 5,
			157, 40, 62, 58, 142, 84, 183, 106, 47, 68, 42, 242, 147, 202,
			67, 195, 249, 77, 155, 4, 130, 156, 223, 180, 73, 32, 200, 249,
			205, 230, 59, 206, 170, 51, 178, 57, 247, 129, 179, 40, 35, 155,
			239, 182, 199, 173, 123, 109, 90, 160, 131, 62, 35, 91, 254, 98,
			51, 41, 133, 253, 189, 94, 153, 183, 23, 93, 185, 212, 19, 250,
			42, 174, 24, 118, 111, 85, 201, 50, 97, 222, 178, 149, 130, 32,
			161, 183, 166, 92, 234, 76, 29, 222, 178, 135, 3, 193, 202, 187,
			101, 15, 7, 107, 189, 123, 151, 126, 91, 38, 139, 48, 178, 237,
			223, 111, 158, 41, 33, 223, 235, 209, 233, 156, 75, 253, 162, 175,
			151, 34, 163, 147, 237, 202, 47, 163, 128, 109, 91, 150, 8, 234,
			100, 123, 218, 37, 140, 212, 25, 217, 110, 56, 218, 17, 202, 200,
			182, 61, 95, 172, 117, 103, 9, 21, 21, 176, 224, 81, 237, 243,
			82, 81, 129, 199, 200, 163, 250, 109, 250, 136, 6, 65, 96, 20,
			245, 216, 159, 110, 110, 188, 162, 162, 76, 183, 168, 42, 69, 89,
			196, 1, 202, 230, 177, 69, 28, 160, 108, 30, 91, 217, 4, 40,
			155, 199, 183, 110, 35, 142, 144, 5, 59, 181, 31, 151, 56, 66,
			143, 145, 157, 250, 52, 42, 59, 52, 56, 118, 223, 160, 178, 67,
			132, 184, 107, 33, 134, 8, 113, 215, 66, 12, 17, 226, 174, 85,
			118, 136, 202, 222, 181, 202, 14, 81, 203, 187, 86, 217, 214, 178,
			141, 77, 104, 24, 183, 247, 90, 141, 77, 136, 50, 220, 171, 224,
			24, 25, 238, 89, 238, 134, 40, 195, 189, 41, 7, 206, 200, 112,
			207, 114, 55, 68, 225, 237, 89, 238, 90, 203, 54, 54, 161, 145,
			225, 254, 127, 209, 216, 132, 40, 168, 253, 10, 148, 113, 111, 223,
			18, 47, 68, 65, 237, 91, 226, 133, 40, 168, 253, 134, 139, 138,
			145, 208, 190, 37, 158, 181, 238, 44, 225, 193, 18, 154, 207, 14,
			253, 181, 55, 121, 176, 132, 40, 152, 67, 91, 211, 66, 20, 204,
			97, 221, 33, 53, 130, 57, 180, 53, 45, 68, 193, 28, 218, 154,
			22, 162, 68, 14, 109, 77, 179, 214, 189, 118, 217, 144, 133, 126,
			192, 200, 179, 255, 141, 134, 44, 52, 194, 33, 207, 170, 196, 24,
			5, 63, 155, 116, 116, 13, 8, 35, 207, 202, 134, 108, 162, 198,
			130, 47, 205, 117, 156, 82, 50, 97, 40, 254, 101, 253, 26, 253,
			157, 71, 131, 9, 188, 107, 61, 247, 63, 105, 254, 218, 3, 247,
			179, 1, 204, 205, 146, 15, 240, 66, 90, 1, 199, 125, 193, 222,
			93, 17, 82, 121, 249, 114, 192, 202, 74, 166, 120, 94, 168, 108,
			204, 43, 188, 47, 86, 89, 235, 73, 53, 136, 176, 149, 17, 25,
			116, 100, 247, 4, 68, 166, 115, 30, 117, 77, 188, 190, 120, 242,
			100, 255, 156, 60, 39, 202, 203, 208, 243, 137, 235, 206, 242, 25,
			121, 126, 99, 222, 89, 132, 145, 231, 75, 31, 211, 191, 149, 158,
			120, 140, 68, 254, 86, 243, 59, 239, 180, 30, 59, 64, 103, 234,
			129, 171, 218, 174, 102, 85, 13, 90, 69, 176, 149, 177, 171, 219,
			197, 51, 184, 186, 244, 182, 160, 236, 208, 78, 87, 24, 255, 41,
			112, 218, 218, 209, 177, 222, 238, 242, 214, 174, 114, 215, 8, 62,
			154, 184, 225, 44, 159, 145, 104, 10, 156, 69, 24, 137, 238, 255,
			128, 254, 217, 71, 119, 125, 70, 250, 254, 71, 205, 63, 249, 14,
			197, 57, 95, 241, 14, 123, 182, 46, 99, 153, 108, 193, 14, 146,
			83, 113, 188, 180, 15, 138, 52, 23, 195, 148, 83, 59, 71, 151,
			31, 86, 137, 31, 47, 176, 139, 199, 137, 136, 19, 72, 240, 47,
			144, 157, 216, 147, 234, 72, 99, 154, 7, 66, 41, 169, 244, 210,
			10, 5, 83, 152, 19, 30, 31, 105, 24, 42, 33, 149, 200, 5,
			31, 139, 76, 44, 179, 158, 232, 187, 191, 40, 67, 169, 181, 232,
			164, 246, 63, 205, 216, 252, 72, 113, 90, 225, 104, 193, 78, 6,
			186, 136, 19, 136, 35, 205, 87, 172, 3, 166, 205, 150, 208, 47,
			34, 21, 101, 57, 231, 80, 226, 147, 25, 135, 99, 145, 166, 230,
			190, 95, 6, 5, 255, 73, 148, 81, 52, 52, 233, 79, 92, 117,
			150, 137, 226, 181, 166, 179, 8, 35, 253, 247, 31, 116, 38, 240,
			15, 208, 131, 127, 5, 0, 0, 255, 255, 223, 229, 133, 116, 137,
			22, 0, 0},
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
	ret, err := discovery.GetDescriptorSet("crrev.Crrev")
	if err != nil {
		panic(err)
	}
	return ret
}
