package ceres.graph

import ceres.AVL.IAVLSet
import ceres.AVL.IAVLTree
import ceres.avl.Serializer

inline data class IRI(val iri: String): Comparable<IRI> {
    override fun compareTo(other: IRI) = iri.compareTo(other.iri)
}

data class Literal(val datatype: String, val value: String): Comparable<Literal> {
    override fun compareTo(other: Literal): Int {
        val x = datatype.compareTo(other.datatype)
        if(x != 0) return x
        return value.compareTo(other.value)
    }
}

data class LangString(val lang: String, val value: String): Comparable<LangString> {
    override fun compareTo(other: LangString): Int {
        val x = lang.compareTo(other.lang)
        if(x != 0) return x
        return value.compareTo(other.value)
    }
}

abstract sealed class RDFValue<T: Comparable<T>>: Comparable<RDFValue<T>> {
    enum class TypeID(val id: Short) {
        IRI(0),
        LITERAL(1),
        STRING(2),
        LANG_STRING(3),
        INT(4),
        LONG(5),
        DOUBLE(6),
        BOOL(7)
    }

    abstract val typeID: TypeID

    abstract val value: T

    override fun compareTo(other: RDFValue<T>): Int {
        val typeCmp = typeID.compareTo(other.typeID)
        if(typeCmp != 0) return typeCmp
        return compare(other)
    }

    fun compare(other: RDFValue<T>): Int = value.compareTo(other.value)

    data class IRIValue(override val value: IRI): RDFValue<IRI>() {
        override val typeID =  TypeID.IRI
    }

    data class LiteralValue(override val value: Literal): RDFValue<Literal>() {
        override val typeID =  TypeID.LITERAL
    }

    data class LangStringValue(override val value: LangString): RDFValue<LangString>() {
        override val typeID = TypeID.LANG_STRING
    }

    data class IntValue(override val value: Int): RDFValue<Int>() {
        override val typeID = TypeID.INT
    }

    data class LongValue(override val value: Long): RDFValue<Long>() {
        override val typeID = TypeID.LONG
    }

    data class DoubleValue(val x: Int)
    data class StringValue(val x: Int)
}



data class EAV(val ent: IRI, val attr: IRI, val value: RDFValue<*>): Comparable<EAV> {
    override fun compareTo(other: EAV): Int {
        val x = ent.compareTo(other.ent)
        if(x != 0) return x
        val y = attr.compareTo(other.attr)
        if(y != 0) return y
        return value.compareTo(other.value)
    }
}

data class VAE(val value: IRI, val attr: IRI, val ent: IRI): Comparable<VAE> {
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
