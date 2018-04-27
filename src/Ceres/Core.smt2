(declare-sort Obj)
(declare-sort IRI)
(declare-sort Set)
(declare-sort Fun)
(declare-datatypes
 ()
 ((Val (AnObj (obj Obj))
        (ANum (num Real))
        (AnInt (int Int))
        (ABool (bool Bool))
        (AString (str String))
        (AnIRI (iri IRI))
        (ASet (set Set))
        (AFun (fun Fun))
        None)
  ;; Expressions need to have only one level of nesting because we can't/don't want to use recursion for instantiation
  (IntExpr
   TheInt
   (Int_i (int_i Int)))
  (AddIntExpr
   (_IntExpr (_intExpr IntExpr))
   (Add_i (add_ia IntExpr) (add_ib IntExpr)))
  (CompIntExpr
   (_AddIntExpr (_addIntExpr AddIntExpr))
   (Gt_i (gt_ia IntExpr) (gt_ib IntExpr))
   (Gte_i (gte_ia IntExpr) (gte_ib IntExpr))
   (Lt_i (lt_ia IntExpr) (lt_ib IntExpr))
   (Lte_i (lte_ia IntExpr) (lte_ib IntExpr)))
  (AndIntExpr
   (_CompIntExpr (_compIntExpr CompIntExpr))
   (And_i (and_ia CompIntExpr) (and_ib CompIntExpr)))
  (OrIntExpr
   (_AndIntExpr (_andIntExpr AndIntExpr))
   (Or_i (or_ia AndIntExpr) (or_ib AndIntExpr)))))

(declare-fun obj_get (Obj String) Val)

(define-fun is-Type ((x Val)) Bool
  (and
   (is-AnObj x)
   (is-ABool (obj_get (obj x) "canBeNone"))
   (is-ABool (obj_get (obj x) "canBeBool"))))

(define-fun app-int-expr ((it Int) (expr IntExpr)) Int
  (match expr
         ((TheInt it)
          ((Int_i x) x))))

;; (declare-fun app-bool-int-expr (Int BoolIntExpr) Bool)

;; (define-fun app-bool-int-expr ((it Int) (expr BoolIntExpr)) Bool
;;   (match expr
;;          (((And_e a b) (and (app-bool-int-expr it a) (app-bool-int-expr it b)))
;;           ((Gt_e a b) (> (app-int-expr it a) (app-int-expr it b)))
;;           ((Lt_e a b) (< (app-int-expr it a) (app-int-expr it b))))))

(define-fun isa ((val Val) (type Val)) Bool
  (or
   (and
    (is-None val)
    (bool (obj_get (obj type) "canBeNone")))
   (and
    (is-ABool val)
    (bool (obj_get (obj type) "canBeBool")))
   false))

(declare-const boolT Val)

(assert (and
         (is-AnObj boolT)
         (is-ABool (obj_get (obj boolT) "canBeBool"))
         (bool (obj_get (obj boolT) "canBeBool"))
         (is-ABool (obj_get (obj boolT) "canBeNone"))
         (not (bool (obj_get (obj boolT) "canBeNone")))))

(push)

(assert (is-Type boolT))

(check-sat)

(pop)

(push)

(assert (not (is-Type boolT)))

(check-sat)

(pop)

(push)

(declare-const x0 Val)

(assert (is-ABool x0))

(push)

(assert (isa x0 boolT))

(check-sat)

(pop)

(push)

(assert (not (isa x0 boolT)))

(check-sat)

(pop)

(pop)

