(declare-sort Obj)
(declare-sort IRI)
(declare-sort Set)
(declare-datatypes
 ()
 ((Term (ObjT (objt Obj))
        (NumT (numt Real))
        (IntT (intt Int))
        (BoolT (boolt Bool))
        (StringT (strt String))
        (SetT (sett Set))
        None)))

(declare-const x () (Set Int))
