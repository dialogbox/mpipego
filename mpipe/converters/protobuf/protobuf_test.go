package protobuf

import (
	"testing"
	"time"

	"github.com/dialogbox/mpipego/common"
)

var protobufData = [][]byte{
	[]byte{0xa, 0xe, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x12, 0xf, 0xa, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3, 0x64, 0x62, 0x30, 0x12, 0x11, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x9, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x36, 0x33, 0x38, 0x31, 0x12, 0x1a, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0xd, 0xa, 0x7, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x74, 0x6c, 0x12, 0x2, 0x18, 0x0, 0x1a, 0xd, 0xa, 0x7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0xa, 0xa, 0x4, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2, 0x18, 0x1, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
	[]byte{0xa, 0x5, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x11, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x9, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x36, 0x33, 0x38, 0x31, 0x12, 0x1a, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1d, 0xa, 0x17, 0x61, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1d, 0xa, 0x10, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x65, 0x61, 0x6b, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x1c, 0xf8, 0x43, 0x41, 0x1a, 0x13, 0xa, 0x6, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x20, 0x54, 0xa8, 0x6a, 0x41, 0x1a, 0x12, 0xa, 0xc, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x25, 0xa, 0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x88, 0xf4, 0xae, 0x8b, 0x41, 0x1a, 0x16, 0xa, 0x10, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1e, 0xa, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x75, 0x66, 0x12, 0x2, 0x18, 0x5, 0x1a, 0x4c, 0xa, 0x6, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x30, 0x12, 0x42, 0xa, 0x40, 0x69, 0x70, 0x3d, 0x31, 0x37, 0x32, 0x2e, 0x33, 0x30, 0x2e, 0x32, 0x31, 0x38, 0x2e, 0x31, 0x34, 0x30, 0x2c, 0x70, 0x6f, 0x72, 0x74, 0x3d, 0x36, 0x33, 0x38, 0x31, 0x2c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x3d, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2c, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3d, 0x31, 0x39, 0x37, 0x32, 0x35, 0x35, 0x30, 0x38, 0x32, 0x37, 0x2c, 0x6c, 0x61, 0x67, 0x3d, 0x31, 0x1a, 0x19, 0xa, 0xc, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x73, 0x79, 0x73, 0x12, 0x9, 0x11, 0x14, 0xae, 0x47, 0xe1, 0xba, 0xa1, 0xb6, 0x40, 0x1a, 0x25, 0xa, 0x18, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6b, 0x62, 0x70, 0x73, 0x12, 0x9, 0x11, 0xb8, 0x1e, 0x85, 0xeb, 0x51, 0xb8, 0xce, 0x3f, 0x1a, 0xd, 0xa, 0x7, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2, 0x18, 0x5, 0x1a, 0x21, 0xa, 0x1b, 0x72, 0x64, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x11, 0xa, 0xb, 0x61, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x15, 0xa, 0xf, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x17, 0xa, 0xf, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x75, 0x61, 0x12, 0x4, 0x18, 0x80, 0xa8, 0x2, 0x1a, 0x24, 0xa, 0x17, 0x6d, 0x65, 0x6d, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x9, 0x11, 0x1f, 0x85, 0xeb, 0x51, 0xb8, 0x1e, 0xfd, 0x3f, 0x1a, 0x21, 0xa, 0x10, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xd, 0xa, 0xb, 0x61, 0x6c, 0x6c, 0x6b, 0x65, 0x79, 0x73, 0x2d, 0x6c, 0x72, 0x75, 0x1a, 0x16, 0xa, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x73, 0x12, 0x2, 0x18, 0x1, 0x1a, 0x1a, 0xa, 0xd, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x9, 0x11, 0xc3, 0xf5, 0x28, 0x5c, 0x8f, 0xd3, 0xab, 0x40, 0x1a, 0x22, 0xa, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x4, 0x18, 0xd6, 0x9b, 0x33, 0x1a, 0x1f, 0xa, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x9, 0x11, 0x0, 0x0, 0xc0, 0x72, 0xaf, 0x64, 0xdd, 0x41, 0x1a, 0x21, 0xa, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6c, 0x65, 0x6e, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x41, 0x1a, 0x1c, 0xa, 0x16, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x18, 0xa, 0xb, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x9c, 0x94, 0x41, 0x41, 0x1a, 0x19, 0xa, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2, 0x18, 0x1, 0x1a, 0x1b, 0xa, 0x15, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x73, 0x79, 0x73, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x13, 0xa, 0xd, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x16, 0xa, 0x9, 0x6d, 0x61, 0x78, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xc0, 0xb, 0x5a, 0xe6, 0x41, 0x1a, 0x15, 0xa, 0xf, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1c, 0xa, 0xf, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x73, 0x73, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf0, 0x4f, 0x41, 0x1a, 0x27, 0xa, 0x1a, 0x72, 0x64, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x99, 0x4f, 0x60, 0x41, 0x1a, 0x28, 0xa, 0x19, 0x61, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x12, 0xb, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0x1a, 0x2a, 0xa, 0x1b, 0x72, 0x64, 0x62, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x67, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x12, 0xb, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0x1a, 0x1f, 0xa, 0x12, 0x72, 0x64, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x9, 0x11, 0x0, 0x0, 0x80, 0x3d, 0x63, 0x1d, 0xd6, 0x41, 0x1a, 0x1e, 0xa, 0x18, 0x72, 0x64, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x67, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1f, 0xa, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x12, 0x2, 0x18, 0x4, 0x1a, 0x1c, 0xa, 0x16, 0x72, 0x64, 0x62, 0x5f, 0x62, 0x67, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1e, 0xa, 0x16, 0x72, 0x64, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x67, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4, 0xa, 0x2, 0x6f, 0x6b, 0x1a, 0xd, 0xa, 0x7, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x15, 0xa, 0xf, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x23, 0xa, 0x16, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x9, 0x11, 0x7b, 0x14, 0xae, 0x47, 0xe1, 0x7a, 0x84, 0x3f, 0x1a, 0x1a, 0xa, 0x14, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x2b, 0xa, 0x1c, 0x61, 0x6f, 0x66, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x12, 0xb, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0x1a, 0x21, 0xa, 0x19, 0x61, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x67, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4, 0xa, 0x2, 0x6f, 0x6b, 0x1a, 0x23, 0xa, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x9, 0x11, 0x0, 0x0, 0xe0, 0xb8, 0x47, 0xe3, 0xa, 0x42, 0x1a, 0x20, 0xa, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1e, 0xa, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x30, 0x41, 0x1a, 0x22, 0xa, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x9, 0x11, 0x0, 0x0, 0xc0, 0x99, 0x8e, 0x26, 0xe6, 0x41, 0x1a, 0x16, 0xa, 0x10, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x16, 0xa, 0x9, 0x6c, 0x72, 0x75, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xa0, 0x37, 0x1, 0x6f, 0x41, 0x1a, 0x2b, 0xa, 0x1e, 0x72, 0x65, 0x70, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x73, 0xaf, 0x60, 0xdd, 0x41, 0x1a, 0x12, 0xa, 0xc, 0x65, 0x76, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x1b, 0xa, 0x15, 0x61, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x20, 0xa, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xc0, 0x16, 0x51, 0x1f, 0x42, 0x1a, 0x17, 0xa, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x12, 0x3, 0x18, 0xb8, 0x1, 0x1a, 0x15, 0xa, 0xf, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x2, 0x18, 0x1, 0x1a, 0xf, 0xa, 0x9, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x12, 0x2, 0x18, 0x1, 0x1a, 0x26, 0xa, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6b, 0x62, 0x70, 0x73, 0x12, 0x9, 0x11, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xe1, 0x3f, 0x1a, 0x1d, 0xa, 0x15, 0x61, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4, 0xa, 0x2, 0x6f, 0x6b, 0x1a, 0x15, 0xa, 0xf, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x15, 0xa, 0xf, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6f, 0x6b, 0x12, 0x2, 0x18, 0x0, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
	[]byte{0xa, 0xe, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0xf, 0xa, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3, 0x64, 0x62, 0x30, 0x12, 0x11, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x9, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x38, 0x30, 0x30, 0x31, 0x12, 0x1a, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0xd, 0xa, 0x7, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x74, 0x6c, 0x12, 0x2, 0x18, 0x0, 0x1a, 0xd, 0xa, 0x7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2, 0x18, 0x0, 0x1a, 0x11, 0xa, 0x4, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xf0, 0x8, 0x74, 0x73, 0x41, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
	[]byte{0xa, 0xe, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0xf, 0xa, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3, 0x64, 0x62, 0x30, 0x12, 0x11, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x9, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x37, 0x30, 0x30, 0x30, 0x12, 0x19, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x5, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0xd, 0xa, 0x7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2, 0x18, 0xb, 0x1a, 0xb, 0xa, 0x4, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x3, 0x18, 0xe2, 0x18, 0x1a, 0x14, 0xa, 0x7, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x74, 0x6c, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xdb, 0x91, 0xcd, 0xd0, 0x41, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
	[]byte{0xa, 0xe, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x12, 0xf, 0xa, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3, 0x64, 0x62, 0x30, 0x12, 0x11, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x9, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x37, 0x30, 0x30, 0x31, 0x12, 0x1a, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0xd, 0xa, 0x7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2, 0x18, 0xa, 0x1a, 0xb, 0xa, 0x4, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x3, 0x18, 0x9e, 0x18, 0x1a, 0x14, 0xa, 0x7, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x74, 0x6c, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0x85, 0x95, 0x3d, 0xc5, 0x41, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
	[]byte{0xa, 0xe, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0xf, 0xa, 0x8, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3, 0x64, 0x62, 0x30, 0x12, 0x16, 0xa, 0x4, 0x68, 0x6f, 0x73, 0x74, 0x12, 0xe, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x41, 0x6d, 0x65, 0x62, 0x61, 0x30, 0x31, 0x67, 0x6d, 0x12, 0xc, 0xa, 0x4, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4, 0x36, 0x33, 0x38, 0x32, 0x12, 0x1a, 0xa, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x6, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x13, 0xa, 0x6, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x9, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0xa, 0x7, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x74, 0x6c, 0x12, 0x9, 0x11, 0x0, 0x0, 0x0, 0xf9, 0xa6, 0x3a, 0xb6, 0x41, 0x1a, 0xd, 0xa, 0x7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2, 0x18, 0x1a, 0x1a, 0xa, 0xa, 0x4, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2, 0x18, 0x1d, 0x20, 0x80, 0xd8, 0xef, 0xef, 0x85, 0xb5, 0xbd, 0xdb, 0x14},
}

var expected = []*common.Metric{
	&common.Metric{Name: "redis_keyspace", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"database": "db0", "host": "Ameba01gm", "port": "6381", "replication_role": "master", "server": "localhost"}, Fields: map[string]interface{}{"avg_ttl": 0, "expires": 0, "keys": 1}},
	&common.Metric{Name: "redis", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"server": "localhost", "host": "Ameba01gm", "port": "6381", "replication_role": "master"}, Fields: map[string]interface{}{"migrate_cached_sockets": 0, "rdb_bgsave_in_progress": 0, "repl_backlog_active": 1, "repl_backlog_histlen": 1.048576e+06, "slave0": "ip=172.30.218.140,port=6381,state=online,offset=1972550827,lag=1", "sync_partial_ok": 0, "keyspace_hits": 0, "total_net_output_bytes": 1.4435284764e+10, "aof_rewrite_scheduled": 0, "sync_partial_err": 0, "used_memory_peak": 2.6174e+06, "connected_slaves": 1, "uptime": 1.3976225e+07, "aof_last_bgrewrite_status": "ok", "aof_last_rewrite_time_sec": -1, "rdb_last_save_time": 1.484098806e+09, "used_cpu_sys_children": 0, "aof_current_rewrite_time_sec": -1, "cluster_enabled": 0, "loading": 0, "pubsub_channels": 1, "used_cpu_user": 3561.78, "used_memory_lua": 37888, "client_longest_output_list": 0, "pubsub_patterns": 0, "sync_full": 1, "used_memory_rss": 4.186112e+06, "expired_keys": 0, "keyspace_misses": 0, "latest_fork_usec": 184, "maxmemory_policy": "allkeys-lru", "mem_fragmentation_ratio": 1.82, "aof_rewrite_in_progress": 0, "master_repl_offset": 1.972551115e+09, "rdb_changes_since_last_save": 0, "rdb_last_save_time_elapsed": 8.551624e+06, "rejected_connections": 0, "repl_backlog_size": 1.048576e+06, "instantaneous_output_kbps": 0.55, "total_connections_received": 839126, "lru_clock": 1.6255421e+07, "total_commands_processed": 5.8056337e+07, "total_net_input_bytes": 2.973005006e+09, "rdb_last_bgsave_status": "ok", "evicted_keys": 0, "rdb_current_bgsave_time_sec": -1, "total_system_memory": 3.3626107904e+10, "used_cpu_sys": 5793.73, "clients": 5, "client_biggest_input_buf": 5, "instantaneous_ops_per_sec": 4, "maxmemory": 3e+09, "repl_backlog_first_byte_offset": 1.97150254e+09, "used_cpu_user_children": 0.01, "aof_enabled": 0, "aof_last_write_status": "ok", "rdb_last_bgsave_time_sec": 0, "keyspace_hitrate": 0, "instantaneous_input_kbps": 0.24, "used_memory": 2.304312e+06, "blocked_clients": 0}},
	&common.Metric{Name: "redis_keyspace", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"database": "db0", "host": "Ameba01gm", "port": "8001", "replication_role": "master", "server": "localhost"}, Fields: map[string]interface{}{"avg_ttl": 0, "expires": 0, "keys": 2.0398223e+07}},
	&common.Metric{Name: "redis_keyspace", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"database": "db0", "host": "Ameba01gm", "port": "7000", "replication_role": "slave", "server": "localhost"}, Fields: map[string]interface{}{"avg_ttl": 1.1276307e+09, "expires": 11, "keys": 3170}},
	&common.Metric{Name: "redis_keyspace", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"database": "db0", "host": "Ameba01gm", "port": "7001", "replication_role": "master", "server": "localhost"}, Fields: map[string]interface{}{"keys": 3102, "avg_ttl": 7.12715018e+08, "expires": 10}},
	&common.Metric{Name: "redis_keyspace", Timestamp: time.Unix(1492650430, 0), Tags: map[string]string{"database": "db0", "host": "corp\\Ameba01gm", "port": "6382", "replication_role": "master", "server": "localhost"}, Fields: map[string]interface{}{"avg_ttl": 3.72942585e+08, "expires": 26, "keys": 29}},
}

func TestProtobufEncoder(t *testing.T) {
	converter := &conv{}

	for _, om := range expected {
		b, err := converter.Encode(om)
		if err != nil {
			t.Error(err)
		}

		m, err := converter.Convert(b)
		if err != nil {
			t.Error(err)
		}

		if !m.Identical(om) {
			t.Logf("Expected:\n%#v\nGot:\n%#v", om, m)
			// fmt.Printf("%#v\n", b)
		}
	}
}

func TestProtobufConverter(t *testing.T) {
	converter := NewConverter()

	for i := range protobufData {
		m, err := converter.Convert(protobufData[i])
		if err != nil {
			t.Error(err)
			continue
		}

		if !m.Identical(expected[i]) {
			t.Logf("Expected:\n%#v\nGot:\n%#v", expected[i], m)
		}
	}
}