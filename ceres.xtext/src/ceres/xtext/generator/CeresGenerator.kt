package ceres.xtext.generator

import org.eclipse.emf.ecore.resource.Resource
import org.eclipse.xtext.generator.AbstractGenerator
import org.eclipse.xtext.generator.IFileSystemAccess2
import org.eclipse.xtext.generator.IGeneratorContext

/**
 * Generates code from your model files on save.
 * 
 * See https://www.eclipse.org/Xtext/documentation/303_runtime_concepts.html#code-generation
 */
open class CeresGeneratorKt: AbstractGenerator() {

	override fun doGenerate(resource: Resource, fsa: IFileSystemAccess2, context: IGeneratorContext ) {
//		fsa.generateFile("test1.kt", "package test1")
		fsa.generateFile("greetings.txt", resource.uri.toString() + " from kt!")
	}
}
