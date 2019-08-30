// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	flowcontrol "k8s.io/kubernetes/pkg/apis/flowcontrol"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowDistinguisherMethod)(nil), (*flowcontrol.FlowDistinguisherMethod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowDistinguisherMethod_To_flowcontrol_FlowDistinguisherMethod(a.(*v1alpha1.FlowDistinguisherMethod), b.(*flowcontrol.FlowDistinguisherMethod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowDistinguisherMethod)(nil), (*v1alpha1.FlowDistinguisherMethod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowDistinguisherMethod_To_v1alpha1_FlowDistinguisherMethod(a.(*flowcontrol.FlowDistinguisherMethod), b.(*v1alpha1.FlowDistinguisherMethod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowSchema)(nil), (*flowcontrol.FlowSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowSchema_To_flowcontrol_FlowSchema(a.(*v1alpha1.FlowSchema), b.(*flowcontrol.FlowSchema), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowSchema)(nil), (*v1alpha1.FlowSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowSchema_To_v1alpha1_FlowSchema(a.(*flowcontrol.FlowSchema), b.(*v1alpha1.FlowSchema), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowSchemaCondition)(nil), (*flowcontrol.FlowSchemaCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowSchemaCondition_To_flowcontrol_FlowSchemaCondition(a.(*v1alpha1.FlowSchemaCondition), b.(*flowcontrol.FlowSchemaCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowSchemaCondition)(nil), (*v1alpha1.FlowSchemaCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowSchemaCondition_To_v1alpha1_FlowSchemaCondition(a.(*flowcontrol.FlowSchemaCondition), b.(*v1alpha1.FlowSchemaCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowSchemaList)(nil), (*flowcontrol.FlowSchemaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowSchemaList_To_flowcontrol_FlowSchemaList(a.(*v1alpha1.FlowSchemaList), b.(*flowcontrol.FlowSchemaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowSchemaList)(nil), (*v1alpha1.FlowSchemaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowSchemaList_To_v1alpha1_FlowSchemaList(a.(*flowcontrol.FlowSchemaList), b.(*v1alpha1.FlowSchemaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowSchemaSpec)(nil), (*flowcontrol.FlowSchemaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec(a.(*v1alpha1.FlowSchemaSpec), b.(*flowcontrol.FlowSchemaSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowSchemaSpec)(nil), (*v1alpha1.FlowSchemaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec(a.(*flowcontrol.FlowSchemaSpec), b.(*v1alpha1.FlowSchemaSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FlowSchemaStatus)(nil), (*flowcontrol.FlowSchemaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus(a.(*v1alpha1.FlowSchemaStatus), b.(*flowcontrol.FlowSchemaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.FlowSchemaStatus)(nil), (*v1alpha1.FlowSchemaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus(a.(*flowcontrol.FlowSchemaStatus), b.(*v1alpha1.FlowSchemaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PolicyRule)(nil), (*flowcontrol.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule(a.(*v1alpha1.PolicyRule), b.(*flowcontrol.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PolicyRule)(nil), (*v1alpha1.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule(a.(*flowcontrol.PolicyRule), b.(*v1alpha1.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PolicyRuleWithSubjects)(nil), (*flowcontrol.PolicyRuleWithSubjects)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PolicyRuleWithSubjects_To_flowcontrol_PolicyRuleWithSubjects(a.(*v1alpha1.PolicyRuleWithSubjects), b.(*flowcontrol.PolicyRuleWithSubjects), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PolicyRuleWithSubjects)(nil), (*v1alpha1.PolicyRuleWithSubjects)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PolicyRuleWithSubjects_To_v1alpha1_PolicyRuleWithSubjects(a.(*flowcontrol.PolicyRuleWithSubjects), b.(*v1alpha1.PolicyRuleWithSubjects), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfiguration)(nil), (*flowcontrol.PriorityLevelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfiguration_To_flowcontrol_PriorityLevelConfiguration(a.(*v1alpha1.PriorityLevelConfiguration), b.(*flowcontrol.PriorityLevelConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfiguration)(nil), (*v1alpha1.PriorityLevelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfiguration_To_v1alpha1_PriorityLevelConfiguration(a.(*flowcontrol.PriorityLevelConfiguration), b.(*v1alpha1.PriorityLevelConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfigurationCondition)(nil), (*flowcontrol.PriorityLevelConfigurationCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfigurationCondition_To_flowcontrol_PriorityLevelConfigurationCondition(a.(*v1alpha1.PriorityLevelConfigurationCondition), b.(*flowcontrol.PriorityLevelConfigurationCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfigurationCondition)(nil), (*v1alpha1.PriorityLevelConfigurationCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfigurationCondition_To_v1alpha1_PriorityLevelConfigurationCondition(a.(*flowcontrol.PriorityLevelConfigurationCondition), b.(*v1alpha1.PriorityLevelConfigurationCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfigurationList)(nil), (*flowcontrol.PriorityLevelConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfigurationList_To_flowcontrol_PriorityLevelConfigurationList(a.(*v1alpha1.PriorityLevelConfigurationList), b.(*flowcontrol.PriorityLevelConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfigurationList)(nil), (*v1alpha1.PriorityLevelConfigurationList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfigurationList_To_v1alpha1_PriorityLevelConfigurationList(a.(*flowcontrol.PriorityLevelConfigurationList), b.(*v1alpha1.PriorityLevelConfigurationList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfigurationReference)(nil), (*flowcontrol.PriorityLevelConfigurationReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference(a.(*v1alpha1.PriorityLevelConfigurationReference), b.(*flowcontrol.PriorityLevelConfigurationReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfigurationReference)(nil), (*v1alpha1.PriorityLevelConfigurationReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference(a.(*flowcontrol.PriorityLevelConfigurationReference), b.(*v1alpha1.PriorityLevelConfigurationReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfigurationSpec)(nil), (*flowcontrol.PriorityLevelConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec(a.(*v1alpha1.PriorityLevelConfigurationSpec), b.(*flowcontrol.PriorityLevelConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfigurationSpec)(nil), (*v1alpha1.PriorityLevelConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec(a.(*flowcontrol.PriorityLevelConfigurationSpec), b.(*v1alpha1.PriorityLevelConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.PriorityLevelConfigurationStatus)(nil), (*flowcontrol.PriorityLevelConfigurationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus(a.(*v1alpha1.PriorityLevelConfigurationStatus), b.(*flowcontrol.PriorityLevelConfigurationStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.PriorityLevelConfigurationStatus)(nil), (*v1alpha1.PriorityLevelConfigurationStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus(a.(*flowcontrol.PriorityLevelConfigurationStatus), b.(*v1alpha1.PriorityLevelConfigurationStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Subject)(nil), (*flowcontrol.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Subject_To_flowcontrol_Subject(a.(*v1alpha1.Subject), b.(*flowcontrol.Subject), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*flowcontrol.Subject)(nil), (*v1alpha1.Subject)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_flowcontrol_Subject_To_v1alpha1_Subject(a.(*flowcontrol.Subject), b.(*v1alpha1.Subject), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_FlowDistinguisherMethod_To_flowcontrol_FlowDistinguisherMethod(in *v1alpha1.FlowDistinguisherMethod, out *flowcontrol.FlowDistinguisherMethod, s conversion.Scope) error {
	out.Type = flowcontrol.FlowDistinguisherMethodType(in.Type)
	return nil
}

// Convert_v1alpha1_FlowDistinguisherMethod_To_flowcontrol_FlowDistinguisherMethod is an autogenerated conversion function.
func Convert_v1alpha1_FlowDistinguisherMethod_To_flowcontrol_FlowDistinguisherMethod(in *v1alpha1.FlowDistinguisherMethod, out *flowcontrol.FlowDistinguisherMethod, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowDistinguisherMethod_To_flowcontrol_FlowDistinguisherMethod(in, out, s)
}

func autoConvert_flowcontrol_FlowDistinguisherMethod_To_v1alpha1_FlowDistinguisherMethod(in *flowcontrol.FlowDistinguisherMethod, out *v1alpha1.FlowDistinguisherMethod, s conversion.Scope) error {
	out.Type = v1alpha1.FlowDistinguisherMethodType(in.Type)
	return nil
}

// Convert_flowcontrol_FlowDistinguisherMethod_To_v1alpha1_FlowDistinguisherMethod is an autogenerated conversion function.
func Convert_flowcontrol_FlowDistinguisherMethod_To_v1alpha1_FlowDistinguisherMethod(in *flowcontrol.FlowDistinguisherMethod, out *v1alpha1.FlowDistinguisherMethod, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowDistinguisherMethod_To_v1alpha1_FlowDistinguisherMethod(in, out, s)
}

func autoConvert_v1alpha1_FlowSchema_To_flowcontrol_FlowSchema(in *v1alpha1.FlowSchema, out *flowcontrol.FlowSchema, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_FlowSchema_To_flowcontrol_FlowSchema is an autogenerated conversion function.
func Convert_v1alpha1_FlowSchema_To_flowcontrol_FlowSchema(in *v1alpha1.FlowSchema, out *flowcontrol.FlowSchema, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowSchema_To_flowcontrol_FlowSchema(in, out, s)
}

func autoConvert_flowcontrol_FlowSchema_To_v1alpha1_FlowSchema(in *flowcontrol.FlowSchema, out *v1alpha1.FlowSchema, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_flowcontrol_FlowSchema_To_v1alpha1_FlowSchema is an autogenerated conversion function.
func Convert_flowcontrol_FlowSchema_To_v1alpha1_FlowSchema(in *flowcontrol.FlowSchema, out *v1alpha1.FlowSchema, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowSchema_To_v1alpha1_FlowSchema(in, out, s)
}

func autoConvert_v1alpha1_FlowSchemaCondition_To_flowcontrol_FlowSchemaCondition(in *v1alpha1.FlowSchemaCondition, out *flowcontrol.FlowSchemaCondition, s conversion.Scope) error {
	out.Type = flowcontrol.FlowSchemaConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_FlowSchemaCondition_To_flowcontrol_FlowSchemaCondition is an autogenerated conversion function.
func Convert_v1alpha1_FlowSchemaCondition_To_flowcontrol_FlowSchemaCondition(in *v1alpha1.FlowSchemaCondition, out *flowcontrol.FlowSchemaCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowSchemaCondition_To_flowcontrol_FlowSchemaCondition(in, out, s)
}

func autoConvert_flowcontrol_FlowSchemaCondition_To_v1alpha1_FlowSchemaCondition(in *flowcontrol.FlowSchemaCondition, out *v1alpha1.FlowSchemaCondition, s conversion.Scope) error {
	out.Type = v1alpha1.FlowSchemaConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_flowcontrol_FlowSchemaCondition_To_v1alpha1_FlowSchemaCondition is an autogenerated conversion function.
func Convert_flowcontrol_FlowSchemaCondition_To_v1alpha1_FlowSchemaCondition(in *flowcontrol.FlowSchemaCondition, out *v1alpha1.FlowSchemaCondition, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowSchemaCondition_To_v1alpha1_FlowSchemaCondition(in, out, s)
}

func autoConvert_v1alpha1_FlowSchemaList_To_flowcontrol_FlowSchemaList(in *v1alpha1.FlowSchemaList, out *flowcontrol.FlowSchemaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]flowcontrol.FlowSchema)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_FlowSchemaList_To_flowcontrol_FlowSchemaList is an autogenerated conversion function.
func Convert_v1alpha1_FlowSchemaList_To_flowcontrol_FlowSchemaList(in *v1alpha1.FlowSchemaList, out *flowcontrol.FlowSchemaList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowSchemaList_To_flowcontrol_FlowSchemaList(in, out, s)
}

func autoConvert_flowcontrol_FlowSchemaList_To_v1alpha1_FlowSchemaList(in *flowcontrol.FlowSchemaList, out *v1alpha1.FlowSchemaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.FlowSchema)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_flowcontrol_FlowSchemaList_To_v1alpha1_FlowSchemaList is an autogenerated conversion function.
func Convert_flowcontrol_FlowSchemaList_To_v1alpha1_FlowSchemaList(in *flowcontrol.FlowSchemaList, out *v1alpha1.FlowSchemaList, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowSchemaList_To_v1alpha1_FlowSchemaList(in, out, s)
}

func autoConvert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec(in *v1alpha1.FlowSchemaSpec, out *flowcontrol.FlowSchemaSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference(&in.PriorityLevelConfiguration, &out.PriorityLevelConfiguration, s); err != nil {
		return err
	}
	out.MatchingPrecedence = in.MatchingPrecedence
	out.DistinguisherMethod = (*flowcontrol.FlowDistinguisherMethod)(unsafe.Pointer(in.DistinguisherMethod))
	out.Rules = *(*[]flowcontrol.PolicyRuleWithSubjects)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec is an autogenerated conversion function.
func Convert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec(in *v1alpha1.FlowSchemaSpec, out *flowcontrol.FlowSchemaSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowSchemaSpec_To_flowcontrol_FlowSchemaSpec(in, out, s)
}

func autoConvert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec(in *flowcontrol.FlowSchemaSpec, out *v1alpha1.FlowSchemaSpec, s conversion.Scope) error {
	if err := Convert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference(&in.PriorityLevelConfiguration, &out.PriorityLevelConfiguration, s); err != nil {
		return err
	}
	out.MatchingPrecedence = in.MatchingPrecedence
	out.DistinguisherMethod = (*v1alpha1.FlowDistinguisherMethod)(unsafe.Pointer(in.DistinguisherMethod))
	out.Rules = *(*[]v1alpha1.PolicyRuleWithSubjects)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec is an autogenerated conversion function.
func Convert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec(in *flowcontrol.FlowSchemaSpec, out *v1alpha1.FlowSchemaSpec, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowSchemaSpec_To_v1alpha1_FlowSchemaSpec(in, out, s)
}

func autoConvert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus(in *v1alpha1.FlowSchemaStatus, out *flowcontrol.FlowSchemaStatus, s conversion.Scope) error {
	out.Conditions = *(*[]flowcontrol.FlowSchemaCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus is an autogenerated conversion function.
func Convert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus(in *v1alpha1.FlowSchemaStatus, out *flowcontrol.FlowSchemaStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlowSchemaStatus_To_flowcontrol_FlowSchemaStatus(in, out, s)
}

func autoConvert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus(in *flowcontrol.FlowSchemaStatus, out *v1alpha1.FlowSchemaStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1alpha1.FlowSchemaCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus is an autogenerated conversion function.
func Convert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus(in *flowcontrol.FlowSchemaStatus, out *v1alpha1.FlowSchemaStatus, s conversion.Scope) error {
	return autoConvert_flowcontrol_FlowSchemaStatus_To_v1alpha1_FlowSchemaStatus(in, out, s)
}

func autoConvert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule(in *v1alpha1.PolicyRule, out *flowcontrol.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule is an autogenerated conversion function.
func Convert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule(in *v1alpha1.PolicyRule, out *flowcontrol.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule(in, out, s)
}

func autoConvert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule(in *flowcontrol.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule is an autogenerated conversion function.
func Convert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule(in *flowcontrol.PolicyRule, out *v1alpha1.PolicyRule, s conversion.Scope) error {
	return autoConvert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule(in, out, s)
}

func autoConvert_v1alpha1_PolicyRuleWithSubjects_To_flowcontrol_PolicyRuleWithSubjects(in *v1alpha1.PolicyRuleWithSubjects, out *flowcontrol.PolicyRuleWithSubjects, s conversion.Scope) error {
	out.Subjects = *(*[]flowcontrol.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_v1alpha1_PolicyRule_To_flowcontrol_PolicyRule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PolicyRuleWithSubjects_To_flowcontrol_PolicyRuleWithSubjects is an autogenerated conversion function.
func Convert_v1alpha1_PolicyRuleWithSubjects_To_flowcontrol_PolicyRuleWithSubjects(in *v1alpha1.PolicyRuleWithSubjects, out *flowcontrol.PolicyRuleWithSubjects, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyRuleWithSubjects_To_flowcontrol_PolicyRuleWithSubjects(in, out, s)
}

func autoConvert_flowcontrol_PolicyRuleWithSubjects_To_v1alpha1_PolicyRuleWithSubjects(in *flowcontrol.PolicyRuleWithSubjects, out *v1alpha1.PolicyRuleWithSubjects, s conversion.Scope) error {
	out.Subjects = *(*[]v1alpha1.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_flowcontrol_PolicyRule_To_v1alpha1_PolicyRule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_flowcontrol_PolicyRuleWithSubjects_To_v1alpha1_PolicyRuleWithSubjects is an autogenerated conversion function.
func Convert_flowcontrol_PolicyRuleWithSubjects_To_v1alpha1_PolicyRuleWithSubjects(in *flowcontrol.PolicyRuleWithSubjects, out *v1alpha1.PolicyRuleWithSubjects, s conversion.Scope) error {
	return autoConvert_flowcontrol_PolicyRuleWithSubjects_To_v1alpha1_PolicyRuleWithSubjects(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfiguration_To_flowcontrol_PriorityLevelConfiguration(in *v1alpha1.PriorityLevelConfiguration, out *flowcontrol.PriorityLevelConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PriorityLevelConfiguration_To_flowcontrol_PriorityLevelConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfiguration_To_flowcontrol_PriorityLevelConfiguration(in *v1alpha1.PriorityLevelConfiguration, out *flowcontrol.PriorityLevelConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfiguration_To_flowcontrol_PriorityLevelConfiguration(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfiguration_To_v1alpha1_PriorityLevelConfiguration(in *flowcontrol.PriorityLevelConfiguration, out *v1alpha1.PriorityLevelConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_flowcontrol_PriorityLevelConfiguration_To_v1alpha1_PriorityLevelConfiguration is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfiguration_To_v1alpha1_PriorityLevelConfiguration(in *flowcontrol.PriorityLevelConfiguration, out *v1alpha1.PriorityLevelConfiguration, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfiguration_To_v1alpha1_PriorityLevelConfiguration(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfigurationCondition_To_flowcontrol_PriorityLevelConfigurationCondition(in *v1alpha1.PriorityLevelConfigurationCondition, out *flowcontrol.PriorityLevelConfigurationCondition, s conversion.Scope) error {
	out.Type = flowcontrol.PriorityLevelConfigurationConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_PriorityLevelConfigurationCondition_To_flowcontrol_PriorityLevelConfigurationCondition is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfigurationCondition_To_flowcontrol_PriorityLevelConfigurationCondition(in *v1alpha1.PriorityLevelConfigurationCondition, out *flowcontrol.PriorityLevelConfigurationCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfigurationCondition_To_flowcontrol_PriorityLevelConfigurationCondition(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfigurationCondition_To_v1alpha1_PriorityLevelConfigurationCondition(in *flowcontrol.PriorityLevelConfigurationCondition, out *v1alpha1.PriorityLevelConfigurationCondition, s conversion.Scope) error {
	out.Type = v1alpha1.PriorityLevelConfigurationConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_flowcontrol_PriorityLevelConfigurationCondition_To_v1alpha1_PriorityLevelConfigurationCondition is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfigurationCondition_To_v1alpha1_PriorityLevelConfigurationCondition(in *flowcontrol.PriorityLevelConfigurationCondition, out *v1alpha1.PriorityLevelConfigurationCondition, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfigurationCondition_To_v1alpha1_PriorityLevelConfigurationCondition(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfigurationList_To_flowcontrol_PriorityLevelConfigurationList(in *v1alpha1.PriorityLevelConfigurationList, out *flowcontrol.PriorityLevelConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]flowcontrol.PriorityLevelConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PriorityLevelConfigurationList_To_flowcontrol_PriorityLevelConfigurationList is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfigurationList_To_flowcontrol_PriorityLevelConfigurationList(in *v1alpha1.PriorityLevelConfigurationList, out *flowcontrol.PriorityLevelConfigurationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfigurationList_To_flowcontrol_PriorityLevelConfigurationList(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfigurationList_To_v1alpha1_PriorityLevelConfigurationList(in *flowcontrol.PriorityLevelConfigurationList, out *v1alpha1.PriorityLevelConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.PriorityLevelConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_flowcontrol_PriorityLevelConfigurationList_To_v1alpha1_PriorityLevelConfigurationList is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfigurationList_To_v1alpha1_PriorityLevelConfigurationList(in *flowcontrol.PriorityLevelConfigurationList, out *v1alpha1.PriorityLevelConfigurationList, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfigurationList_To_v1alpha1_PriorityLevelConfigurationList(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference(in *v1alpha1.PriorityLevelConfigurationReference, out *flowcontrol.PriorityLevelConfigurationReference, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference(in *v1alpha1.PriorityLevelConfigurationReference, out *flowcontrol.PriorityLevelConfigurationReference, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfigurationReference_To_flowcontrol_PriorityLevelConfigurationReference(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference(in *flowcontrol.PriorityLevelConfigurationReference, out *v1alpha1.PriorityLevelConfigurationReference, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference(in *flowcontrol.PriorityLevelConfigurationReference, out *v1alpha1.PriorityLevelConfigurationReference, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfigurationReference_To_v1alpha1_PriorityLevelConfigurationReference(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec(in *v1alpha1.PriorityLevelConfigurationSpec, out *flowcontrol.PriorityLevelConfigurationSpec, s conversion.Scope) error {
	out.AssuredConcurrencyShares = in.AssuredConcurrencyShares
	out.Queues = in.Queues
	out.HandSize = in.HandSize
	out.QueueLengthLimit = in.QueueLengthLimit
	out.Exempt = in.Exempt
	return nil
}

// Convert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec(in *v1alpha1.PriorityLevelConfigurationSpec, out *flowcontrol.PriorityLevelConfigurationSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfigurationSpec_To_flowcontrol_PriorityLevelConfigurationSpec(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec(in *flowcontrol.PriorityLevelConfigurationSpec, out *v1alpha1.PriorityLevelConfigurationSpec, s conversion.Scope) error {
	out.AssuredConcurrencyShares = in.AssuredConcurrencyShares
	out.Queues = in.Queues
	out.HandSize = in.HandSize
	out.QueueLengthLimit = in.QueueLengthLimit
	out.Exempt = in.Exempt
	return nil
}

// Convert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec(in *flowcontrol.PriorityLevelConfigurationSpec, out *v1alpha1.PriorityLevelConfigurationSpec, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfigurationSpec_To_v1alpha1_PriorityLevelConfigurationSpec(in, out, s)
}

func autoConvert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus(in *v1alpha1.PriorityLevelConfigurationStatus, out *flowcontrol.PriorityLevelConfigurationStatus, s conversion.Scope) error {
	out.Conditions = *(*[]flowcontrol.PriorityLevelConfigurationCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus is an autogenerated conversion function.
func Convert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus(in *v1alpha1.PriorityLevelConfigurationStatus, out *flowcontrol.PriorityLevelConfigurationStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityLevelConfigurationStatus_To_flowcontrol_PriorityLevelConfigurationStatus(in, out, s)
}

func autoConvert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus(in *flowcontrol.PriorityLevelConfigurationStatus, out *v1alpha1.PriorityLevelConfigurationStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1alpha1.PriorityLevelConfigurationCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus is an autogenerated conversion function.
func Convert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus(in *flowcontrol.PriorityLevelConfigurationStatus, out *v1alpha1.PriorityLevelConfigurationStatus, s conversion.Scope) error {
	return autoConvert_flowcontrol_PriorityLevelConfigurationStatus_To_v1alpha1_PriorityLevelConfigurationStatus(in, out, s)
}

func autoConvert_v1alpha1_Subject_To_flowcontrol_Subject(in *v1alpha1.Subject, out *flowcontrol.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIGroup = in.APIGroup
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1alpha1_Subject_To_flowcontrol_Subject is an autogenerated conversion function.
func Convert_v1alpha1_Subject_To_flowcontrol_Subject(in *v1alpha1.Subject, out *flowcontrol.Subject, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subject_To_flowcontrol_Subject(in, out, s)
}

func autoConvert_flowcontrol_Subject_To_v1alpha1_Subject(in *flowcontrol.Subject, out *v1alpha1.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIGroup = in.APIGroup
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_flowcontrol_Subject_To_v1alpha1_Subject is an autogenerated conversion function.
func Convert_flowcontrol_Subject_To_v1alpha1_Subject(in *flowcontrol.Subject, out *v1alpha1.Subject, s conversion.Scope) error {
	return autoConvert_flowcontrol_Subject_To_v1alpha1_Subject(in, out, s)
}