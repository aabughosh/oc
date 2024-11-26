// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	securityv1 "github.com/openshift/api/security/v1"
	internal "github.com/openshift/client-go/security/applyconfigurations/internal"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// SecurityContextConstraintsApplyConfiguration represents a declarative configuration of the SecurityContextConstraints type for use
// with apply.
type SecurityContextConstraintsApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Priority                         *int32                                               `json:"priority,omitempty"`
	AllowPrivilegedContainer         *bool                                                `json:"allowPrivilegedContainer,omitempty"`
	DefaultAddCapabilities           []corev1.Capability                                  `json:"defaultAddCapabilities,omitempty"`
	RequiredDropCapabilities         []corev1.Capability                                  `json:"requiredDropCapabilities,omitempty"`
	AllowedCapabilities              []corev1.Capability                                  `json:"allowedCapabilities,omitempty"`
	AllowHostDirVolumePlugin         *bool                                                `json:"allowHostDirVolumePlugin,omitempty"`
	Volumes                          []securityv1.FSType                                  `json:"volumes,omitempty"`
	AllowedFlexVolumes               []AllowedFlexVolumeApplyConfiguration                `json:"allowedFlexVolumes,omitempty"`
	AllowHostNetwork                 *bool                                                `json:"allowHostNetwork,omitempty"`
	AllowHostPorts                   *bool                                                `json:"allowHostPorts,omitempty"`
	AllowHostPID                     *bool                                                `json:"allowHostPID,omitempty"`
	AllowHostIPC                     *bool                                                `json:"allowHostIPC,omitempty"`
	UserNamespaceLevel               *securityv1.NamespaceLevelType                       `json:"userNamespaceLevel,omitempty"`
	DefaultAllowPrivilegeEscalation  *bool                                                `json:"defaultAllowPrivilegeEscalation,omitempty"`
	AllowPrivilegeEscalation         *bool                                                `json:"allowPrivilegeEscalation,omitempty"`
	SELinuxContext                   *SELinuxContextStrategyOptionsApplyConfiguration     `json:"seLinuxContext,omitempty"`
	RunAsUser                        *RunAsUserStrategyOptionsApplyConfiguration          `json:"runAsUser,omitempty"`
	SupplementalGroups               *SupplementalGroupsStrategyOptionsApplyConfiguration `json:"supplementalGroups,omitempty"`
	FSGroup                          *FSGroupStrategyOptionsApplyConfiguration            `json:"fsGroup,omitempty"`
	ReadOnlyRootFilesystem           *bool                                                `json:"readOnlyRootFilesystem,omitempty"`
	Users                            []string                                             `json:"users,omitempty"`
	Groups                           []string                                             `json:"groups,omitempty"`
	SeccompProfiles                  []string                                             `json:"seccompProfiles,omitempty"`
	AllowedUnsafeSysctls             []string                                             `json:"allowedUnsafeSysctls,omitempty"`
	ForbiddenSysctls                 []string                                             `json:"forbiddenSysctls,omitempty"`
}

// SecurityContextConstraints constructs a declarative configuration of the SecurityContextConstraints type for use with
// apply.
func SecurityContextConstraints(name string) *SecurityContextConstraintsApplyConfiguration {
	b := &SecurityContextConstraintsApplyConfiguration{}
	b.WithName(name)
	b.WithKind("SecurityContextConstraints")
	b.WithAPIVersion("security.openshift.io/v1")
	return b
}

// ExtractSecurityContextConstraints extracts the applied configuration owned by fieldManager from
// securityContextConstraints. If no managedFields are found in securityContextConstraints for fieldManager, a
// SecurityContextConstraintsApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// securityContextConstraints must be a unmodified SecurityContextConstraints API object that was retrieved from the Kubernetes API.
// ExtractSecurityContextConstraints provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractSecurityContextConstraints(securityContextConstraints *securityv1.SecurityContextConstraints, fieldManager string) (*SecurityContextConstraintsApplyConfiguration, error) {
	return extractSecurityContextConstraints(securityContextConstraints, fieldManager, "")
}

// ExtractSecurityContextConstraintsStatus is the same as ExtractSecurityContextConstraints except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractSecurityContextConstraintsStatus(securityContextConstraints *securityv1.SecurityContextConstraints, fieldManager string) (*SecurityContextConstraintsApplyConfiguration, error) {
	return extractSecurityContextConstraints(securityContextConstraints, fieldManager, "status")
}

func extractSecurityContextConstraints(securityContextConstraints *securityv1.SecurityContextConstraints, fieldManager string, subresource string) (*SecurityContextConstraintsApplyConfiguration, error) {
	b := &SecurityContextConstraintsApplyConfiguration{}
	err := managedfields.ExtractInto(securityContextConstraints, internal.Parser().Type("com.github.openshift.api.security.v1.SecurityContextConstraints"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(securityContextConstraints.Name)

	b.WithKind("SecurityContextConstraints")
	b.WithAPIVersion("security.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithKind(value string) *SecurityContextConstraintsApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAPIVersion(value string) *SecurityContextConstraintsApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithName(value string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithGenerateName(value string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithNamespace(value string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithUID(value types.UID) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithResourceVersion(value string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithGeneration(value int64) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithCreationTimestamp(value metav1.Time) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *SecurityContextConstraintsApplyConfiguration) WithLabels(entries map[string]string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *SecurityContextConstraintsApplyConfiguration) WithAnnotations(entries map[string]string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *SecurityContextConstraintsApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *SecurityContextConstraintsApplyConfiguration) WithFinalizers(values ...string) *SecurityContextConstraintsApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *SecurityContextConstraintsApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithPriority(value int32) *SecurityContextConstraintsApplyConfiguration {
	b.Priority = &value
	return b
}

// WithAllowPrivilegedContainer sets the AllowPrivilegedContainer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowPrivilegedContainer field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowPrivilegedContainer(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowPrivilegedContainer = &value
	return b
}

// WithDefaultAddCapabilities adds the given value to the DefaultAddCapabilities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DefaultAddCapabilities field.
func (b *SecurityContextConstraintsApplyConfiguration) WithDefaultAddCapabilities(values ...corev1.Capability) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.DefaultAddCapabilities = append(b.DefaultAddCapabilities, values[i])
	}
	return b
}

// WithRequiredDropCapabilities adds the given value to the RequiredDropCapabilities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RequiredDropCapabilities field.
func (b *SecurityContextConstraintsApplyConfiguration) WithRequiredDropCapabilities(values ...corev1.Capability) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.RequiredDropCapabilities = append(b.RequiredDropCapabilities, values[i])
	}
	return b
}

// WithAllowedCapabilities adds the given value to the AllowedCapabilities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedCapabilities field.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowedCapabilities(values ...corev1.Capability) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.AllowedCapabilities = append(b.AllowedCapabilities, values[i])
	}
	return b
}

// WithAllowHostDirVolumePlugin sets the AllowHostDirVolumePlugin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowHostDirVolumePlugin field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowHostDirVolumePlugin(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowHostDirVolumePlugin = &value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *SecurityContextConstraintsApplyConfiguration) WithVolumes(values ...securityv1.FSType) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}

// WithAllowedFlexVolumes adds the given value to the AllowedFlexVolumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedFlexVolumes field.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowedFlexVolumes(values ...*AllowedFlexVolumeApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAllowedFlexVolumes")
		}
		b.AllowedFlexVolumes = append(b.AllowedFlexVolumes, *values[i])
	}
	return b
}

// WithAllowHostNetwork sets the AllowHostNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowHostNetwork field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowHostNetwork(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowHostNetwork = &value
	return b
}

// WithAllowHostPorts sets the AllowHostPorts field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowHostPorts field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowHostPorts(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowHostPorts = &value
	return b
}

// WithAllowHostPID sets the AllowHostPID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowHostPID field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowHostPID(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowHostPID = &value
	return b
}

// WithAllowHostIPC sets the AllowHostIPC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowHostIPC field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowHostIPC(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowHostIPC = &value
	return b
}

// WithUserNamespaceLevel sets the UserNamespaceLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UserNamespaceLevel field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithUserNamespaceLevel(value securityv1.NamespaceLevelType) *SecurityContextConstraintsApplyConfiguration {
	b.UserNamespaceLevel = &value
	return b
}

// WithDefaultAllowPrivilegeEscalation sets the DefaultAllowPrivilegeEscalation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultAllowPrivilegeEscalation field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithDefaultAllowPrivilegeEscalation(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.DefaultAllowPrivilegeEscalation = &value
	return b
}

// WithAllowPrivilegeEscalation sets the AllowPrivilegeEscalation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowPrivilegeEscalation field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowPrivilegeEscalation(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.AllowPrivilegeEscalation = &value
	return b
}

// WithSELinuxContext sets the SELinuxContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SELinuxContext field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithSELinuxContext(value *SELinuxContextStrategyOptionsApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	b.SELinuxContext = value
	return b
}

// WithRunAsUser sets the RunAsUser field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsUser field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithRunAsUser(value *RunAsUserStrategyOptionsApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	b.RunAsUser = value
	return b
}

// WithSupplementalGroups sets the SupplementalGroups field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SupplementalGroups field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithSupplementalGroups(value *SupplementalGroupsStrategyOptionsApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	b.SupplementalGroups = value
	return b
}

// WithFSGroup sets the FSGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSGroup field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithFSGroup(value *FSGroupStrategyOptionsApplyConfiguration) *SecurityContextConstraintsApplyConfiguration {
	b.FSGroup = value
	return b
}

// WithReadOnlyRootFilesystem sets the ReadOnlyRootFilesystem field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnlyRootFilesystem field is set to the value of the last call.
func (b *SecurityContextConstraintsApplyConfiguration) WithReadOnlyRootFilesystem(value bool) *SecurityContextConstraintsApplyConfiguration {
	b.ReadOnlyRootFilesystem = &value
	return b
}

// WithUsers adds the given value to the Users field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Users field.
func (b *SecurityContextConstraintsApplyConfiguration) WithUsers(values ...string) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.Users = append(b.Users, values[i])
	}
	return b
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *SecurityContextConstraintsApplyConfiguration) WithGroups(values ...string) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.Groups = append(b.Groups, values[i])
	}
	return b
}

// WithSeccompProfiles adds the given value to the SeccompProfiles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SeccompProfiles field.
func (b *SecurityContextConstraintsApplyConfiguration) WithSeccompProfiles(values ...string) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.SeccompProfiles = append(b.SeccompProfiles, values[i])
	}
	return b
}

// WithAllowedUnsafeSysctls adds the given value to the AllowedUnsafeSysctls field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedUnsafeSysctls field.
func (b *SecurityContextConstraintsApplyConfiguration) WithAllowedUnsafeSysctls(values ...string) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.AllowedUnsafeSysctls = append(b.AllowedUnsafeSysctls, values[i])
	}
	return b
}

// WithForbiddenSysctls adds the given value to the ForbiddenSysctls field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ForbiddenSysctls field.
func (b *SecurityContextConstraintsApplyConfiguration) WithForbiddenSysctls(values ...string) *SecurityContextConstraintsApplyConfiguration {
	for i := range values {
		b.ForbiddenSysctls = append(b.ForbiddenSysctls, values[i])
	}
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *SecurityContextConstraintsApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
