package ceres.graph

import ceres.AVL.IAVLSet
import ceres.AVL.IAVLTree

abstract sealed class RDFValue: Comparable<RDFValue> {
    override fun compareTo(other: RDFValue): Int {
        TODO()
    }

    data class IRI(val iri: String): RDFValue() {
//        override fun compareTo(other: RDFValue) =
//            if(other is IRI) iri.compareTo(other.iri)
//            else -1
    }

    data class Literal(val datatype: String, val value: String): RDFValue() {
//        override fun compareTo(other: RDFValue): Int {
//            if(other !is Literal)
//                return 1
//            val x = type.compareTo(other.type)
//            if(x != 0) return x
//            if(langTag != null) {
//                val otherLangTag = other.langTag
//                if(otherLangTag == null)
//                    return 1
//                val y = langTag.compareTo(otherLangTag)
//                if(y != 0) return y
//            }
//            return value.compareTo(other.value)
//        }
    }

    data class LangString(val x: String, val lang: String)
    data class IntValue(val x: Int)
    data class LongValue(val x: Int)
    data class DoubleValue(val x: Int)
    data class StringValue(val x: Int)
}



data class EAV(val ent: RDFValue.IRI, val attr: RDFValue.IRI, val value: RDFValue): Comparable<EAV> {
    override fun compareTo(other: EAV): Int {
        val x = ent.compareTo(other.ent)
        if(x != 0) return x
        val y = attr.compareTo(other.attr)
        if(y != 0) return y
        return value.compareTo(other.value)
    }
}

data class VAE(val value: RDFValue.IRI, val attr: RDFValue.IRI, val ent: RDFValue.IRI): Comparable<VAE> {
    override fun compareTo(other: VAE): Int {
        val x = value.compareTo(other.value)
        if(x != 0) return x
        val y = attr.compareTo(other.attr)
        if(y != 0) return y
        return ent.compareTo(other.ent)
    }
}

data class VE(val value: RDFValue, val ent: RDFValue.IRI): Comparable<VE> {
    override fun compareTo(other: VE): Int {
        val x = value.compareTo(other.value)
        if(x != 0) return x
        return ent.compareTo(other.ent)
    }
}

data class Graph(
        val eav: IAVLSet<EAV>,
        val vae: IAVLSet<VAE>

) {
    fun get(iri: RDFValue.IRI): Node? {
        TODO()
    }
}

interface ValueSet {

}

class Node {
    //fun get(attr: IRI): ValueSet
}
