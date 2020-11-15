package runtime

// Builder builds a new Scheme for mapping go types to Kubernetes GroupVersionKinds.
type Builder struct {
	GroupVersion GroupVersion
	SchemeBuilder
}

// Register adds one or objects to the SchemeBuilder so they can be added to a Scheme.  Register mutates bld.
func (bld *Builder) Register(object ...Object) *Builder {
	bld.SchemeBuilder.Register(func(scheme *Scheme) error {
		scheme.AddSupportType(bld.GroupVersion, object...)
		return nil
	})
	return bld
}

// RegisterAll registers all types from the Builder argument.  RegisterAll mutates bld.
func (bld *Builder) RegisterAll(b *Builder) *Builder {
	bld.SchemeBuilder = append(bld.SchemeBuilder, b.SchemeBuilder...)
	return bld
}

// AddToScheme adds all registered types to s.
func (bld *Builder) AddToScheme(s *Scheme) error {
	return bld.SchemeBuilder.AddToScheme(s)
}

// Build returns a new Scheme containing the registered types.
func (bld *Builder) Build() (*Scheme, error) {
	s := NewScheme()
	return s, bld.AddToScheme(s)
}
