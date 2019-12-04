// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pessoa-service.proto

package pessoa

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Telefone_TelefoneTipo int32

const (
	Telefone_CELULAR  Telefone_TelefoneTipo = 0
	Telefone_CASA     Telefone_TelefoneTipo = 1
	Telefone_TRABALHO Telefone_TelefoneTipo = 2
)

var Telefone_TelefoneTipo_name = map[int32]string{
	0: "CELULAR",
	1: "CASA",
	2: "TRABALHO",
}

var Telefone_TelefoneTipo_value = map[string]int32{
	"CELULAR":  0,
	"CASA":     1,
	"TRABALHO": 2,
}

func (x Telefone_TelefoneTipo) String() string {
	return proto.EnumName(Telefone_TelefoneTipo_name, int32(x))
}

func (Telefone_TelefoneTipo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{0, 0}
}

// Objeto Telefone
type Telefone struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Dd                   string                `protobuf:"bytes,2,opt,name=dd,proto3" json:"dd,omitempty"`
	Numero               string                `protobuf:"bytes,3,opt,name=numero,proto3" json:"numero,omitempty"`
	Tipo                 Telefone_TelefoneTipo `protobuf:"varint,4,opt,name=tipo,proto3,enum=pessoa.Telefone_TelefoneTipo" json:"tipo,omitempty"`
	Reminder             *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Telefone) Reset()         { *m = Telefone{} }
func (m *Telefone) String() string { return proto.CompactTextString(m) }
func (*Telefone) ProtoMessage()    {}
func (*Telefone) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{0}
}

func (m *Telefone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telefone.Unmarshal(m, b)
}
func (m *Telefone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telefone.Marshal(b, m, deterministic)
}
func (m *Telefone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telefone.Merge(m, src)
}
func (m *Telefone) XXX_Size() int {
	return xxx_messageInfo_Telefone.Size(m)
}
func (m *Telefone) XXX_DiscardUnknown() {
	xxx_messageInfo_Telefone.DiscardUnknown(m)
}

var xxx_messageInfo_Telefone proto.InternalMessageInfo

func (m *Telefone) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Telefone) GetDd() string {
	if m != nil {
		return m.Dd
	}
	return ""
}

func (m *Telefone) GetNumero() string {
	if m != nil {
		return m.Numero
	}
	return ""
}

func (m *Telefone) GetTipo() Telefone_TelefoneTipo {
	if m != nil {
		return m.Tipo
	}
	return Telefone_CELULAR
}

func (m *Telefone) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// objeto identificador
type Identificador struct {
	Cpf                  int64    `protobuf:"varint,1,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Rg                   int64    `protobuf:"varint,2,opt,name=rg,proto3" json:"rg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identificador) Reset()         { *m = Identificador{} }
func (m *Identificador) String() string { return proto.CompactTextString(m) }
func (*Identificador) ProtoMessage()    {}
func (*Identificador) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{1}
}

func (m *Identificador) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identificador.Unmarshal(m, b)
}
func (m *Identificador) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identificador.Marshal(b, m, deterministic)
}
func (m *Identificador) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identificador.Merge(m, src)
}
func (m *Identificador) XXX_Size() int {
	return xxx_messageInfo_Identificador.Size(m)
}
func (m *Identificador) XXX_DiscardUnknown() {
	xxx_messageInfo_Identificador.DiscardUnknown(m)
}

var xxx_messageInfo_Identificador proto.InternalMessageInfo

func (m *Identificador) GetCpf() int64 {
	if m != nil {
		return m.Cpf
	}
	return 0
}

func (m *Identificador) GetRg() int64 {
	if m != nil {
		return m.Rg
	}
	return 0
}

// Objeto Pessoa
type Pessoa struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nome                 string               `protobuf:"bytes,2,opt,name=nome,proto3" json:"nome,omitempty"`
	DtNascimento         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=dtNascimento,proto3" json:"dtNascimento,omitempty"`
	Identificador        []*Identificador     `protobuf:"bytes,4,rep,name=identificador,proto3" json:"identificador,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Telefone             []*Telefone          `protobuf:"bytes,6,rep,name=telefone,proto3" json:"telefone,omitempty"`
	Endereco             []*Endereco          `protobuf:"bytes,7,rep,name=endereco,proto3" json:"endereco,omitempty"`
	Reminder             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pessoa) Reset()         { *m = Pessoa{} }
func (m *Pessoa) String() string { return proto.CompactTextString(m) }
func (*Pessoa) ProtoMessage()    {}
func (*Pessoa) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{2}
}

func (m *Pessoa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pessoa.Unmarshal(m, b)
}
func (m *Pessoa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pessoa.Marshal(b, m, deterministic)
}
func (m *Pessoa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pessoa.Merge(m, src)
}
func (m *Pessoa) XXX_Size() int {
	return xxx_messageInfo_Pessoa.Size(m)
}
func (m *Pessoa) XXX_DiscardUnknown() {
	xxx_messageInfo_Pessoa.DiscardUnknown(m)
}

var xxx_messageInfo_Pessoa proto.InternalMessageInfo

func (m *Pessoa) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pessoa) GetNome() string {
	if m != nil {
		return m.Nome
	}
	return ""
}

func (m *Pessoa) GetDtNascimento() *timestamp.Timestamp {
	if m != nil {
		return m.DtNascimento
	}
	return nil
}

func (m *Pessoa) GetIdentificador() []*Identificador {
	if m != nil {
		return m.Identificador
	}
	return nil
}

func (m *Pessoa) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Pessoa) GetTelefone() []*Telefone {
	if m != nil {
		return m.Telefone
	}
	return nil
}

func (m *Pessoa) GetEndereco() []*Endereco {
	if m != nil {
		return m.Endereco
	}
	return nil
}

func (m *Pessoa) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Requisicao para criar uma nova pessoa
type CriarPessoaRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Pessoa               *Pessoa  `protobuf:"bytes,2,opt,name=pessoa,proto3" json:"pessoa,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriarPessoaRequest) Reset()         { *m = CriarPessoaRequest{} }
func (m *CriarPessoaRequest) String() string { return proto.CompactTextString(m) }
func (*CriarPessoaRequest) ProtoMessage()    {}
func (*CriarPessoaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{3}
}

func (m *CriarPessoaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriarPessoaRequest.Unmarshal(m, b)
}
func (m *CriarPessoaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriarPessoaRequest.Marshal(b, m, deterministic)
}
func (m *CriarPessoaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriarPessoaRequest.Merge(m, src)
}
func (m *CriarPessoaRequest) XXX_Size() int {
	return xxx_messageInfo_CriarPessoaRequest.Size(m)
}
func (m *CriarPessoaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CriarPessoaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CriarPessoaRequest proto.InternalMessageInfo

func (m *CriarPessoaRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CriarPessoaRequest) GetPessoa() *Pessoa {
	if m != nil {
		return m.Pessoa
	}
	return nil
}

// Resposta com ID da pessoa Criada
type CriarPessoaResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriarPessoaResponse) Reset()         { *m = CriarPessoaResponse{} }
func (m *CriarPessoaResponse) String() string { return proto.CompactTextString(m) }
func (*CriarPessoaResponse) ProtoMessage()    {}
func (*CriarPessoaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{4}
}

func (m *CriarPessoaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriarPessoaResponse.Unmarshal(m, b)
}
func (m *CriarPessoaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriarPessoaResponse.Marshal(b, m, deterministic)
}
func (m *CriarPessoaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriarPessoaResponse.Merge(m, src)
}
func (m *CriarPessoaResponse) XXX_Size() int {
	return xxx_messageInfo_CriarPessoaResponse.Size(m)
}
func (m *CriarPessoaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CriarPessoaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CriarPessoaResponse proto.InternalMessageInfo

func (m *CriarPessoaResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CriarPessoaResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

//CriarIdentificadorResponse cria um identificador para a pessoa response
type CriarIdentificadorResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Cpf                  int64    `protobuf:"varint,2,opt,name=cpf,proto3" json:"cpf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriarIdentificadorResponse) Reset()         { *m = CriarIdentificadorResponse{} }
func (m *CriarIdentificadorResponse) String() string { return proto.CompactTextString(m) }
func (*CriarIdentificadorResponse) ProtoMessage()    {}
func (*CriarIdentificadorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{5}
}

func (m *CriarIdentificadorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriarIdentificadorResponse.Unmarshal(m, b)
}
func (m *CriarIdentificadorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriarIdentificadorResponse.Marshal(b, m, deterministic)
}
func (m *CriarIdentificadorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriarIdentificadorResponse.Merge(m, src)
}
func (m *CriarIdentificadorResponse) XXX_Size() int {
	return xxx_messageInfo_CriarIdentificadorResponse.Size(m)
}
func (m *CriarIdentificadorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CriarIdentificadorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CriarIdentificadorResponse proto.InternalMessageInfo

func (m *CriarIdentificadorResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CriarIdentificadorResponse) GetCpf() int64 {
	if m != nil {
		return m.Cpf
	}
	return 0
}

//CriarIdentificadorRequest cria um identificador para a pessoa request
type CriarIdentificadorRequest struct {
	Identificador        *Identificador `protobuf:"bytes,1,opt,name=identificador,proto3" json:"identificador,omitempty"`
	Api                  string         `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CriarIdentificadorRequest) Reset()         { *m = CriarIdentificadorRequest{} }
func (m *CriarIdentificadorRequest) String() string { return proto.CompactTextString(m) }
func (*CriarIdentificadorRequest) ProtoMessage()    {}
func (*CriarIdentificadorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3daf7753e5a8644f, []int{6}
}

func (m *CriarIdentificadorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriarIdentificadorRequest.Unmarshal(m, b)
}
func (m *CriarIdentificadorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriarIdentificadorRequest.Marshal(b, m, deterministic)
}
func (m *CriarIdentificadorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriarIdentificadorRequest.Merge(m, src)
}
func (m *CriarIdentificadorRequest) XXX_Size() int {
	return xxx_messageInfo_CriarIdentificadorRequest.Size(m)
}
func (m *CriarIdentificadorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CriarIdentificadorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CriarIdentificadorRequest proto.InternalMessageInfo

func (m *CriarIdentificadorRequest) GetIdentificador() *Identificador {
	if m != nil {
		return m.Identificador
	}
	return nil
}

func (m *CriarIdentificadorRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func init() {
	proto.RegisterEnum("pessoa.Telefone_TelefoneTipo", Telefone_TelefoneTipo_name, Telefone_TelefoneTipo_value)
	proto.RegisterType((*Telefone)(nil), "pessoa.Telefone")
	proto.RegisterType((*Identificador)(nil), "pessoa.Identificador")
	proto.RegisterType((*Pessoa)(nil), "pessoa.Pessoa")
	proto.RegisterType((*CriarPessoaRequest)(nil), "pessoa.CriarPessoaRequest")
	proto.RegisterType((*CriarPessoaResponse)(nil), "pessoa.CriarPessoaResponse")
	proto.RegisterType((*CriarIdentificadorResponse)(nil), "pessoa.CriarIdentificadorResponse")
	proto.RegisterType((*CriarIdentificadorRequest)(nil), "pessoa.CriarIdentificadorRequest")
}

func init() { proto.RegisterFile("pessoa-service.proto", fileDescriptor_3daf7753e5a8644f) }

var fileDescriptor_3daf7753e5a8644f = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0xe9, 0xc7, 0xd2, 0xdb, 0x0f, 0x55, 0x66, 0x4c, 0x21, 0x08, 0x51, 0xf2, 0x80, 0xfa,
	0x00, 0x99, 0xd6, 0x49, 0xf0, 0x80, 0x84, 0x56, 0xaa, 0x49, 0x20, 0x55, 0x63, 0xf2, 0xca, 0x13,
	0x4f, 0x59, 0x7c, 0x5b, 0x19, 0x35, 0x71, 0x70, 0x52, 0x7e, 0x16, 0x7f, 0x86, 0x7f, 0xc0, 0x2f,
	0x41, 0xb1, 0x9d, 0x68, 0xa6, 0x65, 0x82, 0x37, 0xfb, 0xe6, 0xdc, 0xe3, 0x7b, 0xce, 0xb9, 0x81,
	0xe3, 0x1c, 0x8b, 0x42, 0xc4, 0xaf, 0x0a, 0x94, 0xdf, 0x79, 0x82, 0x51, 0x2e, 0x45, 0x29, 0x48,
	0x57, 0x57, 0x83, 0x67, 0x1b, 0x21, 0x36, 0x5b, 0x3c, 0x55, 0xd5, 0xdb, 0xdd, 0xfa, 0xb4, 0xe4,
	0x29, 0x16, 0x65, 0x9c, 0xe6, 0x1a, 0x18, 0x9c, 0x60, 0xc6, 0x50, 0x62, 0x22, 0x6c, 0x82, 0xf0,
	0x97, 0x03, 0xde, 0x0a, 0xb7, 0xb8, 0x16, 0x19, 0x92, 0x11, 0xb8, 0x9c, 0xf9, 0xce, 0xc4, 0x99,
	0xb6, 0xa8, 0xcb, 0x59, 0x75, 0x67, 0xcc, 0x77, 0x27, 0xce, 0xb4, 0x47, 0x5d, 0xc6, 0xc8, 0x09,
	0x74, 0xb3, 0x5d, 0x8a, 0x52, 0xf8, 0x2d, 0x55, 0x33, 0x37, 0x72, 0x06, 0xed, 0x92, 0xe7, 0xc2,
	0x6f, 0x4f, 0x9c, 0xe9, 0x68, 0xf6, 0x34, 0xd2, 0x43, 0x45, 0x35, 0x6f, 0x73, 0x58, 0xf1, 0x5c,
	0x50, 0x05, 0x25, 0xaf, 0xc1, 0x93, 0x98, 0xf2, 0x6a, 0x28, 0xbf, 0x33, 0x71, 0xa6, 0xfd, 0x59,
	0x10, 0x69, 0x0d, 0x51, 0xad, 0x21, 0x5a, 0xd5, 0x1a, 0x68, 0x83, 0x0d, 0xcf, 0x61, 0x70, 0x97,
	0x8d, 0xf4, 0xe1, 0x68, 0x71, 0xb9, 0xfc, 0xbc, 0x9c, 0xd3, 0xf1, 0x03, 0xe2, 0x41, 0x7b, 0x31,
	0xbf, 0x99, 0x8f, 0x1d, 0x32, 0x00, 0x6f, 0x45, 0xe7, 0xef, 0xe7, 0xcb, 0x0f, 0x9f, 0xc6, 0x6e,
	0x78, 0x06, 0xc3, 0x8f, 0x0c, 0xb3, 0x92, 0xaf, 0x79, 0x12, 0x33, 0x21, 0xc9, 0x18, 0x5a, 0x49,
	0xbe, 0x36, 0x4a, 0xab, 0x63, 0x25, 0x55, 0x6e, 0x94, 0xd4, 0x16, 0x75, 0xe5, 0x26, 0xfc, 0xe9,
	0x42, 0xf7, 0x5a, 0xc9, 0xd8, 0x73, 0x85, 0x40, 0x3b, 0x13, 0x29, 0x1a, 0x5f, 0xd4, 0x99, 0xbc,
	0x83, 0x01, 0x2b, 0xaf, 0xe2, 0x22, 0xe1, 0x29, 0x66, 0xa5, 0xf6, 0xe7, 0x7e, 0x49, 0x16, 0x9e,
	0xbc, 0x85, 0x21, 0xbf, 0x3b, 0xa1, 0xdf, 0x9e, 0xb4, 0xa6, 0xfd, 0xd9, 0xa3, 0xda, 0x4a, 0x6b,
	0x7c, 0x6a, 0x63, 0xc9, 0x31, 0x74, 0x30, 0x8d, 0xf9, 0x56, 0x19, 0xd9, 0xa3, 0xfa, 0x42, 0x5e,
	0x82, 0x57, 0x1a, 0xa7, 0xfc, 0xae, 0x62, 0x1b, 0xff, 0x19, 0x0c, 0x6d, 0x10, 0x15, 0xba, 0xde,
	0x10, 0xff, 0xc8, 0x46, 0x5f, 0x9a, 0x3a, 0x6d, 0x10, 0x56, 0x7a, 0xde, 0x7f, 0xa4, 0x77, 0x05,
	0x64, 0x21, 0x79, 0x2c, 0xb5, 0xb3, 0x14, 0xbf, 0xed, 0xb0, 0x28, 0xab, 0x34, 0xe2, 0x9c, 0x2b,
	0x87, 0x7b, 0xb4, 0x3a, 0x92, 0x17, 0x60, 0x16, 0x5b, 0x99, 0xdc, 0x9f, 0x8d, 0xea, 0x59, 0x4c,
	0xa3, 0xf9, 0x1a, 0xbe, 0x81, 0x87, 0x16, 0x5f, 0x91, 0x8b, 0xac, 0xc0, 0x03, 0x84, 0x3a, 0x43,
	0xb7, 0xce, 0x30, 0xbc, 0x80, 0x40, 0x35, 0xda, 0xbe, 0xfe, 0xbd, 0xdf, 0x2c, 0x8c, 0xdb, 0x2c,
	0x4c, 0xf8, 0x15, 0x1e, 0x1f, 0x62, 0xd0, 0x8a, 0xf6, 0xe2, 0x74, 0x94, 0x8c, 0x7f, 0x8b, 0xd3,
	0xbc, 0xee, 0x36, 0xaf, 0xcf, 0x7e, 0x38, 0x30, 0xd4, 0x12, 0x6f, 0xf4, 0xcf, 0x4b, 0x2e, 0xa0,
	0x93, 0x54, 0xaf, 0x93, 0xa0, 0xa6, 0xdc, 0xf7, 0x35, 0x78, 0x72, 0xf0, 0x9b, 0xd1, 0xf8, 0xc5,
	0x44, 0x61, 0xff, 0x18, 0xcf, 0xad, 0x96, 0x43, 0xda, 0x82, 0xf0, 0x3e, 0x88, 0x26, 0xbf, 0x76,
	0x6e, 0xbb, 0x6a, 0x0f, 0xce, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x98, 0x01, 0x52, 0x2e, 0xb7,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PessoaServiceClient is the client API for PessoaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PessoaServiceClient interface {
	// Cria uma nova pessoa
	Criar(ctx context.Context, in *CriarPessoaRequest, opts ...grpc.CallOption) (*CriarPessoaResponse, error)
	CriarIdentificador(ctx context.Context, in *CriarIdentificadorRequest, opts ...grpc.CallOption) (*CriarIdentificadorResponse, error)
}

type pessoaServiceClient struct {
	cc *grpc.ClientConn
}

func NewPessoaServiceClient(cc *grpc.ClientConn) PessoaServiceClient {
	return &pessoaServiceClient{cc}
}

func (c *pessoaServiceClient) Criar(ctx context.Context, in *CriarPessoaRequest, opts ...grpc.CallOption) (*CriarPessoaResponse, error) {
	out := new(CriarPessoaResponse)
	err := c.cc.Invoke(ctx, "/pessoa.PessoaService/criar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pessoaServiceClient) CriarIdentificador(ctx context.Context, in *CriarIdentificadorRequest, opts ...grpc.CallOption) (*CriarIdentificadorResponse, error) {
	out := new(CriarIdentificadorResponse)
	err := c.cc.Invoke(ctx, "/pessoa.PessoaService/CriarIdentificador", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PessoaServiceServer is the server API for PessoaService service.
type PessoaServiceServer interface {
	// Cria uma nova pessoa
	Criar(context.Context, *CriarPessoaRequest) (*CriarPessoaResponse, error)
	CriarIdentificador(context.Context, *CriarIdentificadorRequest) (*CriarIdentificadorResponse, error)
}

// UnimplementedPessoaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPessoaServiceServer struct {
}

func (*UnimplementedPessoaServiceServer) Criar(ctx context.Context, req *CriarPessoaRequest) (*CriarPessoaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Criar not implemented")
}
func (*UnimplementedPessoaServiceServer) CriarIdentificador(ctx context.Context, req *CriarIdentificadorRequest) (*CriarIdentificadorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CriarIdentificador not implemented")
}

func RegisterPessoaServiceServer(s *grpc.Server, srv PessoaServiceServer) {
	s.RegisterService(&_PessoaService_serviceDesc, srv)
}

func _PessoaService_Criar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CriarPessoaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PessoaServiceServer).Criar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pessoa.PessoaService/Criar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PessoaServiceServer).Criar(ctx, req.(*CriarPessoaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PessoaService_CriarIdentificador_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CriarIdentificadorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PessoaServiceServer).CriarIdentificador(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pessoa.PessoaService/CriarIdentificador",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PessoaServiceServer).CriarIdentificador(ctx, req.(*CriarIdentificadorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PessoaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pessoa.PessoaService",
	HandlerType: (*PessoaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "criar",
			Handler:    _PessoaService_Criar_Handler,
		},
		{
			MethodName: "CriarIdentificador",
			Handler:    _PessoaService_CriarIdentificador_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pessoa-service.proto",
}
