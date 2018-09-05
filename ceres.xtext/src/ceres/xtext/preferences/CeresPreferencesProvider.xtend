package ceres.xtext.preferences

import org.eclipse.emf.ecore.resource.Resource
import org.eclipse.xtext.preferences.IPreferenceValues
import org.eclipse.xtext.preferences.IPreferenceValuesProvider
import org.eclipse.xtext.preferences.TypedPreferenceValues
import org.eclipse.xtext.preferences.MapBasedPreferenceValues

class CeresPreferencesProvider implements IPreferenceValuesProvider {
	override IPreferenceValues getPreferenceValues(Resource context) {
		return new TypedPreferenceValues(new MapBasedPreferenceValues())
	}
}
