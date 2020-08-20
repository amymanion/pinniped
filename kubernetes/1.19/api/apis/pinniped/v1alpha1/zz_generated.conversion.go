// +build !ignore_autogenerated

/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"

	pinniped "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/pinniped"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CredentialRequest)(nil), (*pinniped.CredentialRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequest_To_pinniped_CredentialRequest(a.(*CredentialRequest), b.(*pinniped.CredentialRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequest)(nil), (*CredentialRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequest_To_v1alpha1_CredentialRequest(a.(*pinniped.CredentialRequest), b.(*CredentialRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CredentialRequestCredential)(nil), (*pinniped.CredentialRequestCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequestCredential_To_pinniped_CredentialRequestCredential(a.(*CredentialRequestCredential), b.(*pinniped.CredentialRequestCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequestCredential)(nil), (*CredentialRequestCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequestCredential_To_v1alpha1_CredentialRequestCredential(a.(*pinniped.CredentialRequestCredential), b.(*CredentialRequestCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CredentialRequestList)(nil), (*pinniped.CredentialRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequestList_To_pinniped_CredentialRequestList(a.(*CredentialRequestList), b.(*pinniped.CredentialRequestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequestList)(nil), (*CredentialRequestList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequestList_To_v1alpha1_CredentialRequestList(a.(*pinniped.CredentialRequestList), b.(*CredentialRequestList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CredentialRequestSpec)(nil), (*pinniped.CredentialRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec(a.(*CredentialRequestSpec), b.(*pinniped.CredentialRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequestSpec)(nil), (*CredentialRequestSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec(a.(*pinniped.CredentialRequestSpec), b.(*CredentialRequestSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CredentialRequestStatus)(nil), (*pinniped.CredentialRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus(a.(*CredentialRequestStatus), b.(*pinniped.CredentialRequestStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequestStatus)(nil), (*CredentialRequestStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus(a.(*pinniped.CredentialRequestStatus), b.(*CredentialRequestStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CredentialRequestTokenCredential)(nil), (*pinniped.CredentialRequestTokenCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CredentialRequestTokenCredential_To_pinniped_CredentialRequestTokenCredential(a.(*CredentialRequestTokenCredential), b.(*pinniped.CredentialRequestTokenCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pinniped.CredentialRequestTokenCredential)(nil), (*CredentialRequestTokenCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_pinniped_CredentialRequestTokenCredential_To_v1alpha1_CredentialRequestTokenCredential(a.(*pinniped.CredentialRequestTokenCredential), b.(*CredentialRequestTokenCredential), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CredentialRequest_To_pinniped_CredentialRequest(in *CredentialRequest, out *pinniped.CredentialRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CredentialRequest_To_pinniped_CredentialRequest is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequest_To_pinniped_CredentialRequest(in *CredentialRequest, out *pinniped.CredentialRequest, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequest_To_pinniped_CredentialRequest(in, out, s)
}

func autoConvert_pinniped_CredentialRequest_To_v1alpha1_CredentialRequest(in *pinniped.CredentialRequest, out *CredentialRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_pinniped_CredentialRequest_To_v1alpha1_CredentialRequest is an autogenerated conversion function.
func Convert_pinniped_CredentialRequest_To_v1alpha1_CredentialRequest(in *pinniped.CredentialRequest, out *CredentialRequest, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequest_To_v1alpha1_CredentialRequest(in, out, s)
}

func autoConvert_v1alpha1_CredentialRequestCredential_To_pinniped_CredentialRequestCredential(in *CredentialRequestCredential, out *pinniped.CredentialRequestCredential, s conversion.Scope) error {
	out.ExpirationTimestamp = in.ExpirationTimestamp
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_v1alpha1_CredentialRequestCredential_To_pinniped_CredentialRequestCredential is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequestCredential_To_pinniped_CredentialRequestCredential(in *CredentialRequestCredential, out *pinniped.CredentialRequestCredential, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequestCredential_To_pinniped_CredentialRequestCredential(in, out, s)
}

func autoConvert_pinniped_CredentialRequestCredential_To_v1alpha1_CredentialRequestCredential(in *pinniped.CredentialRequestCredential, out *CredentialRequestCredential, s conversion.Scope) error {
	out.ExpirationTimestamp = in.ExpirationTimestamp
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_pinniped_CredentialRequestCredential_To_v1alpha1_CredentialRequestCredential is an autogenerated conversion function.
func Convert_pinniped_CredentialRequestCredential_To_v1alpha1_CredentialRequestCredential(in *pinniped.CredentialRequestCredential, out *CredentialRequestCredential, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequestCredential_To_v1alpha1_CredentialRequestCredential(in, out, s)
}

func autoConvert_v1alpha1_CredentialRequestList_To_pinniped_CredentialRequestList(in *CredentialRequestList, out *pinniped.CredentialRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]pinniped.CredentialRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CredentialRequestList_To_pinniped_CredentialRequestList is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequestList_To_pinniped_CredentialRequestList(in *CredentialRequestList, out *pinniped.CredentialRequestList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequestList_To_pinniped_CredentialRequestList(in, out, s)
}

func autoConvert_pinniped_CredentialRequestList_To_v1alpha1_CredentialRequestList(in *pinniped.CredentialRequestList, out *CredentialRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CredentialRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_pinniped_CredentialRequestList_To_v1alpha1_CredentialRequestList is an autogenerated conversion function.
func Convert_pinniped_CredentialRequestList_To_v1alpha1_CredentialRequestList(in *pinniped.CredentialRequestList, out *CredentialRequestList, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequestList_To_v1alpha1_CredentialRequestList(in, out, s)
}

func autoConvert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec(in *CredentialRequestSpec, out *pinniped.CredentialRequestSpec, s conversion.Scope) error {
	out.Type = pinniped.CredentialType(in.Type)
	out.Token = (*pinniped.CredentialRequestTokenCredential)(unsafe.Pointer(in.Token))
	return nil
}

// Convert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec(in *CredentialRequestSpec, out *pinniped.CredentialRequestSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequestSpec_To_pinniped_CredentialRequestSpec(in, out, s)
}

func autoConvert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec(in *pinniped.CredentialRequestSpec, out *CredentialRequestSpec, s conversion.Scope) error {
	out.Type = CredentialType(in.Type)
	out.Token = (*CredentialRequestTokenCredential)(unsafe.Pointer(in.Token))
	return nil
}

// Convert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec is an autogenerated conversion function.
func Convert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec(in *pinniped.CredentialRequestSpec, out *CredentialRequestSpec, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequestSpec_To_v1alpha1_CredentialRequestSpec(in, out, s)
}

func autoConvert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus(in *CredentialRequestStatus, out *pinniped.CredentialRequestStatus, s conversion.Scope) error {
	out.Credential = (*pinniped.CredentialRequestCredential)(unsafe.Pointer(in.Credential))
	out.Message = (*string)(unsafe.Pointer(in.Message))
	return nil
}

// Convert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus(in *CredentialRequestStatus, out *pinniped.CredentialRequestStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequestStatus_To_pinniped_CredentialRequestStatus(in, out, s)
}

func autoConvert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus(in *pinniped.CredentialRequestStatus, out *CredentialRequestStatus, s conversion.Scope) error {
	out.Credential = (*CredentialRequestCredential)(unsafe.Pointer(in.Credential))
	out.Message = (*string)(unsafe.Pointer(in.Message))
	return nil
}

// Convert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus is an autogenerated conversion function.
func Convert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus(in *pinniped.CredentialRequestStatus, out *CredentialRequestStatus, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequestStatus_To_v1alpha1_CredentialRequestStatus(in, out, s)
}

func autoConvert_v1alpha1_CredentialRequestTokenCredential_To_pinniped_CredentialRequestTokenCredential(in *CredentialRequestTokenCredential, out *pinniped.CredentialRequestTokenCredential, s conversion.Scope) error {
	out.Value = in.Value
	return nil
}

// Convert_v1alpha1_CredentialRequestTokenCredential_To_pinniped_CredentialRequestTokenCredential is an autogenerated conversion function.
func Convert_v1alpha1_CredentialRequestTokenCredential_To_pinniped_CredentialRequestTokenCredential(in *CredentialRequestTokenCredential, out *pinniped.CredentialRequestTokenCredential, s conversion.Scope) error {
	return autoConvert_v1alpha1_CredentialRequestTokenCredential_To_pinniped_CredentialRequestTokenCredential(in, out, s)
}

func autoConvert_pinniped_CredentialRequestTokenCredential_To_v1alpha1_CredentialRequestTokenCredential(in *pinniped.CredentialRequestTokenCredential, out *CredentialRequestTokenCredential, s conversion.Scope) error {
	out.Value = in.Value
	return nil
}

// Convert_pinniped_CredentialRequestTokenCredential_To_v1alpha1_CredentialRequestTokenCredential is an autogenerated conversion function.
func Convert_pinniped_CredentialRequestTokenCredential_To_v1alpha1_CredentialRequestTokenCredential(in *pinniped.CredentialRequestTokenCredential, out *CredentialRequestTokenCredential, s conversion.Scope) error {
	return autoConvert_pinniped_CredentialRequestTokenCredential_To_v1alpha1_CredentialRequestTokenCredential(in, out, s)
}
