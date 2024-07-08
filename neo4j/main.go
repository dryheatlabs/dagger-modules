// A module for working with Neo4j
//
// Module for building Neo4J containers, serving or querying a running Neo4J database.
package main

type Neo4j struct {
	User     *Secret
	Password *Secret

	Version string
}

func (n *Neo4j) WithVersion(version string) *Neo4j {
	n.Version = version
	return n
}

// Get the neo4j container with the directory in it
func (n *Neo4j) Container(platform string) *Container {
	opts := ContainerOpts{}
	if platform != "" {
		opts.Platform = Platform(platform)
	}

	return dag.Container(opts).
		From("neo4j" + n.Version).
		WithExposedPort(7474).
		WithExposedPort(7687)
}
