// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/genesis.proto

package v012

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the gitopia module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	ReleaseList       []Release      `protobuf:"bytes,15,rep,name=releaseList,proto3" json:"releaseList"`
	ReleaseCount      uint64         `protobuf:"varint,16,opt,name=releaseCount,proto3" json:"releaseCount,omitempty"`
	PullRequestList   []PullRequest  `protobuf:"bytes,13,rep,name=pullRequestList,proto3" json:"pullRequestList"`
	PullRequestCount  uint64         `protobuf:"varint,14,opt,name=pullRequestCount,proto3" json:"pullRequestCount,omitempty"`
	OrganizationList  []Organization `protobuf:"bytes,11,rep,name=organizationList,proto3" json:"organizationList"`
	OrganizationCount uint64         `protobuf:"varint,12,opt,name=organizationCount,proto3" json:"organizationCount,omitempty"`
	CommentList       []Comment      `protobuf:"bytes,9,rep,name=commentList,proto3" json:"commentList"`
	CommentCount      uint64         `protobuf:"varint,10,opt,name=commentCount,proto3" json:"commentCount,omitempty"`
	IssueList         []Issue        `protobuf:"bytes,7,rep,name=issueList,proto3" json:"issueList"`
	IssueCount        uint64         `protobuf:"varint,8,opt,name=issueCount,proto3" json:"issueCount,omitempty"`
	RepositoryList    []Repository   `protobuf:"bytes,5,rep,name=repositoryList,proto3" json:"repositoryList"`
	RepositoryCount   uint64         `protobuf:"varint,6,opt,name=repositoryCount,proto3" json:"repositoryCount,omitempty"`
	UserList          []User         `protobuf:"bytes,3,rep,name=userList,proto3" json:"userList"`
	UserCount         uint64         `protobuf:"varint,4,opt,name=userCount,proto3" json:"userCount,omitempty"`
	WhoisList         []Whois        `protobuf:"bytes,1,rep,name=whoisList,proto3" json:"whoisList"`
	WhoisCount        uint64         `protobuf:"varint,2,opt,name=whoisCount,proto3" json:"whoisCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe28ed7a80acf9ab, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetReleaseList() []Release {
	if m != nil {
		return m.ReleaseList
	}
	return nil
}

func (m *GenesisState) GetReleaseCount() uint64 {
	if m != nil {
		return m.ReleaseCount
	}
	return 0
}

func (m *GenesisState) GetPullRequestList() []PullRequest {
	if m != nil {
		return m.PullRequestList
	}
	return nil
}

func (m *GenesisState) GetPullRequestCount() uint64 {
	if m != nil {
		return m.PullRequestCount
	}
	return 0
}

func (m *GenesisState) GetOrganizationList() []Organization {
	if m != nil {
		return m.OrganizationList
	}
	return nil
}

func (m *GenesisState) GetOrganizationCount() uint64 {
	if m != nil {
		return m.OrganizationCount
	}
	return 0
}

func (m *GenesisState) GetCommentList() []Comment {
	if m != nil {
		return m.CommentList
	}
	return nil
}

func (m *GenesisState) GetCommentCount() uint64 {
	if m != nil {
		return m.CommentCount
	}
	return 0
}

func (m *GenesisState) GetIssueList() []Issue {
	if m != nil {
		return m.IssueList
	}
	return nil
}

func (m *GenesisState) GetIssueCount() uint64 {
	if m != nil {
		return m.IssueCount
	}
	return 0
}

func (m *GenesisState) GetRepositoryList() []Repository {
	if m != nil {
		return m.RepositoryList
	}
	return nil
}

func (m *GenesisState) GetRepositoryCount() uint64 {
	if m != nil {
		return m.RepositoryCount
	}
	return 0
}

func (m *GenesisState) GetUserList() []User {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *GenesisState) GetUserCount() uint64 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *GenesisState) GetWhoisList() []Whois {
	if m != nil {
		return m.WhoisList
	}
	return nil
}

func (m *GenesisState) GetWhoisCount() uint64 {
	if m != nil {
		return m.WhoisCount
	}
	return 0
}

func init() {
	// proto.RegisterType((*GenesisState)(nil), "gitopia.gitopia.gitopia.GenesisState")
}

func init() { proto.RegisterFile("gitopia/genesis.proto", fileDescriptor_fe28ed7a80acf9ab) }

var fileDescriptor_fe28ed7a80acf9ab = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x13, 0x56, 0xb6, 0xf5, 0xb4, 0xac, 0xc5, 0x80, 0x28, 0x15, 0x98, 0x6a, 0x80, 0x54,
	0x4d, 0x28, 0x95, 0xe0, 0x01, 0x90, 0x3a, 0x24, 0x40, 0x42, 0x02, 0x02, 0x68, 0x12, 0x77, 0xd9,
	0x64, 0x65, 0x96, 0xda, 0x38, 0xc4, 0x8e, 0x60, 0x3c, 0x05, 0x77, 0xbc, 0xd2, 0x2e, 0x77, 0xc9,
	0x15, 0x42, 0xed, 0x8b, 0x20, 0x1f, 0xdb, 0x89, 0x49, 0x95, 0xed, 0x2a, 0xf6, 0x77, 0xfe, 0xfc,
	0x8e, 0xbf, 0x38, 0x81, 0x3b, 0x29, 0x57, 0x22, 0xe7, 0xc9, 0x2c, 0x65, 0x19, 0x93, 0x5c, 0x46,
	0x79, 0x21, 0x94, 0x20, 0x77, 0xad, 0x1c, 0x35, 0x9e, 0xe3, 0xdb, 0xa9, 0x48, 0x05, 0xe6, 0xcc,
	0xf4, 0xca, 0xa4, 0x8f, 0xab, 0x2e, 0x05, 0x5b, 0xb0, 0x44, 0x32, 0x2b, 0xdf, 0x73, 0x72, 0x5e,
	0x2e, 0x16, 0x31, 0xfb, 0x5a, 0x32, 0xa9, 0x6c, 0x68, 0xec, 0x42, 0xa2, 0x48, 0x93, 0x8c, 0xff,
	0x48, 0x14, 0x17, 0x59, 0xb3, 0xdb, 0x89, 0x58, 0x2e, 0x59, 0xe6, 0x4a, 0x6e, 0x39, 0x99, 0x4b,
	0x59, 0x3a, 0xc4, 0xa8, 0x26, 0xe7, 0x42, 0x72, 0x25, 0x8a, 0x33, 0x1b, 0x21, 0x2e, 0x52, 0x4a,
	0x56, 0x34, 0x5b, 0x7c, 0x3b, 0x15, 0xee, 0xac, 0xfb, 0xbf, 0x76, 0xa0, 0xff, 0xca, 0x9c, 0xfe,
	0xa3, 0x4a, 0x14, 0x23, 0xaf, 0xa1, 0x67, 0xcf, 0xf1, 0x96, 0x4b, 0x35, 0x1a, 0x4c, 0xb6, 0xa6,
	0xbd, 0x67, 0x93, 0xa8, 0xc5, 0x92, 0x28, 0x36, 0xb9, 0xf3, 0xce, 0xf9, 0x9f, 0x87, 0x41, 0xec,
	0x97, 0x92, 0x7d, 0xe8, 0xdb, 0xed, 0xa1, 0x28, 0x33, 0x35, 0x1a, 0x4e, 0xc2, 0x69, 0x27, 0xfe,
	0x4f, 0x23, 0x9f, 0x60, 0xe0, 0xd9, 0x83, 0xc4, 0x1b, 0x48, 0x7c, 0xdc, 0x4a, 0x7c, 0x5f, 0xe7,
	0x5b, 0x6a, 0xb3, 0x05, 0x39, 0x80, 0xa1, 0x27, 0x19, 0xfa, 0x1e, 0xd2, 0x37, 0x74, 0x72, 0x04,
	0x43, 0xff, 0x2d, 0xe0, 0x08, 0x3d, 0x1c, 0xe1, 0x49, 0xeb, 0x08, 0xef, 0xbc, 0x02, 0x3b, 0xc3,
	0x46, 0x13, 0xf2, 0x14, 0x6e, 0xfa, 0x9a, 0x99, 0xa2, 0x8f, 0x53, 0x6c, 0x06, 0xb4, 0xed, 0xf6,
	0x85, 0xe3, 0x04, 0xdd, 0x2b, 0x6c, 0x3f, 0x34, 0xb9, 0xce, 0x76, 0xaf, 0x54, 0xdb, 0x6e, 0xb7,
	0x06, 0x09, 0xc6, 0x76, 0x5f, 0x23, 0x73, 0xe8, 0xe2, 0x3d, 0x42, 0xd6, 0x0e, 0xb2, 0x68, 0x2b,
	0xeb, 0x8d, 0xce, 0xb4, 0xa4, 0xba, 0x8c, 0x50, 0x00, 0xdc, 0x18, 0xca, 0x2e, 0x52, 0x3c, 0x85,
	0x7c, 0x80, 0xbd, 0xfa, 0x5a, 0x22, 0xe8, 0x3a, 0x82, 0x1e, 0x5d, 0x72, 0x97, 0x5c, 0xba, 0xa5,
	0x35, 0x1a, 0x90, 0x29, 0x0c, 0x6a, 0xc5, 0x70, 0xb7, 0x91, 0xdb, 0x94, 0xc9, 0x0b, 0xd8, 0xd5,
	0x37, 0x1f, 0xb1, 0x5b, 0x88, 0x7d, 0xd0, 0x8a, 0xfd, 0x2c, 0x59, 0x61, 0x81, 0x55, 0x11, 0xb9,
	0x0f, 0x5d, 0xbd, 0x36, 0x90, 0x0e, 0x42, 0x6a, 0x41, 0xfb, 0x87, 0x1f, 0x11, 0xf6, 0x0f, 0xaf,
	0xf0, 0xef, 0x48, 0x67, 0x3a, 0xff, 0xaa, 0x32, 0xed, 0x1f, 0x6e, 0x0c, 0xe2, 0x9a, 0xf1, 0xaf,
	0x56, 0xe6, 0x2f, 0xcf, 0x57, 0x34, 0xbc, 0x58, 0xd1, 0xf0, 0xef, 0x8a, 0x86, 0x3f, 0xd7, 0x34,
	0xb8, 0x58, 0xd3, 0xe0, 0xf7, 0x9a, 0x06, 0x5f, 0x0e, 0x52, 0xae, 0x4e, 0xcb, 0xe3, 0xe8, 0x44,
	0x2c, 0x67, 0xd5, 0x1f, 0xcc, 0x3e, 0xbf, 0x57, 0x2b, 0x75, 0x96, 0x33, 0x79, 0xbc, 0x8d, 0x9f,
	0xf9, 0xf3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x49, 0xff, 0xfd, 0x2d, 0xeb, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReleaseCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReleaseCount))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.ReleaseList) > 0 {
		for iNdEx := len(m.ReleaseList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReleaseList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.PullRequestCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PullRequestCount))
		i--
		dAtA[i] = 0x70
	}
	if len(m.PullRequestList) > 0 {
		for iNdEx := len(m.PullRequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PullRequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.OrganizationCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.OrganizationCount))
		i--
		dAtA[i] = 0x60
	}
	if len(m.OrganizationList) > 0 {
		for iNdEx := len(m.OrganizationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrganizationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.CommentCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CommentCount))
		i--
		dAtA[i] = 0x50
	}
	if len(m.CommentList) > 0 {
		for iNdEx := len(m.CommentList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommentList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.IssueCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.IssueCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.IssueList) > 0 {
		for iNdEx := len(m.IssueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IssueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.RepositoryCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RepositoryCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RepositoryList) > 0 {
		for iNdEx := len(m.RepositoryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RepositoryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.UserCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UserCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.UserList) > 0 {
		for iNdEx := len(m.UserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.WhoisCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.WhoisCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.WhoisList) > 0 {
		for iNdEx := len(m.WhoisList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhoisList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhoisList) > 0 {
		for _, e := range m.WhoisList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.WhoisCount != 0 {
		n += 1 + sovGenesis(uint64(m.WhoisCount))
	}
	if len(m.UserList) > 0 {
		for _, e := range m.UserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.UserCount != 0 {
		n += 1 + sovGenesis(uint64(m.UserCount))
	}
	if len(m.RepositoryList) > 0 {
		for _, e := range m.RepositoryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.RepositoryCount != 0 {
		n += 1 + sovGenesis(uint64(m.RepositoryCount))
	}
	if len(m.IssueList) > 0 {
		for _, e := range m.IssueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.IssueCount != 0 {
		n += 1 + sovGenesis(uint64(m.IssueCount))
	}
	if len(m.CommentList) > 0 {
		for _, e := range m.CommentList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CommentCount != 0 {
		n += 1 + sovGenesis(uint64(m.CommentCount))
	}
	if len(m.OrganizationList) > 0 {
		for _, e := range m.OrganizationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.OrganizationCount != 0 {
		n += 1 + sovGenesis(uint64(m.OrganizationCount))
	}
	if len(m.PullRequestList) > 0 {
		for _, e := range m.PullRequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PullRequestCount != 0 {
		n += 1 + sovGenesis(uint64(m.PullRequestCount))
	}
	if len(m.ReleaseList) > 0 {
		for _, e := range m.ReleaseList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ReleaseCount != 0 {
		n += 2 + sovGenesis(uint64(m.ReleaseCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhoisList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhoisList = append(m.WhoisList, Whois{})
			if err := m.WhoisList[len(m.WhoisList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhoisCount", wireType)
			}
			m.WhoisCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WhoisCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserList = append(m.UserList, User{})
			if err := m.UserList[len(m.UserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCount", wireType)
			}
			m.UserCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepositoryList = append(m.RepositoryList, Repository{})
			if err := m.RepositoryList[len(m.RepositoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryCount", wireType)
			}
			m.RepositoryCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepositoryCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssueList = append(m.IssueList, Issue{})
			if err := m.IssueList[len(m.IssueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueCount", wireType)
			}
			m.IssueCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IssueCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommentList = append(m.CommentList, Comment{})
			if err := m.CommentList[len(m.CommentList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentCount", wireType)
			}
			m.CommentCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommentCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationList = append(m.OrganizationList, Organization{})
			if err := m.OrganizationList[len(m.OrganizationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationCount", wireType)
			}
			m.OrganizationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrganizationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PullRequestList = append(m.PullRequestList, PullRequest{})
			if err := m.PullRequestList[len(m.PullRequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequestCount", wireType)
			}
			m.PullRequestCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PullRequestCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleaseList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReleaseList = append(m.ReleaseList, Release{})
			if err := m.ReleaseList[len(m.ReleaseList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleaseCount", wireType)
			}
			m.ReleaseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReleaseCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
