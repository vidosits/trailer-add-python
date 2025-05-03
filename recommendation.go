package recommendation

import (
	"encoding/json"
)

func environmentHasPython(environment Environment) bool {
	for _, pkg := range environment.Packages {
		if pkg.Name == "python" {
			return true
		}
	}
	return false
}

func Match(serializedImage []byte) (bool, error) {
	var image Image

	if err := json.Unmarshal(serializedImage, &image); err != nil {
		return false, err
	}

	// Check if ANY environment is missing a conda package named "python"
	for _, env := range image.Configuration.Environments {
		if len(env.Packages) > 0 && !environmentHasPython(env) {
			return true, nil
		}
	}

	// we got this far, so all environments have python, this recommendation is not applicable
	return false, nil
}

func Recommend(serializedImage []byte) ([]byte, error) {
	var image Image

	if err := json.Unmarshal(serializedImage, &image); err != nil {
		return nil, err
	}

	// Add an unconstrained version of python to each environment
	for envName, env := range image.Configuration.Environments {
		if environmentHasPython(env) || len(env.Packages) == 0 {
			continue
		}

		pythonPackage := Package[Version]{
			Name: "python",
			Version: Version{
				Specifier:  UnconstrainedVersion,
				Constraint: "*",
			},
			Channel: "conda-forge",
		}

		// add the package to the environment
		env.Packages = append(env.Packages, pythonPackage)

		// save the environment back to the config
		image.Configuration.Environments[envName] = env

	}

	// Convert back to JSON
	updatedImage, err := json.Marshal(image)
	if err != nil {
		return nil, err
	}

	return updatedImage, nil
}
