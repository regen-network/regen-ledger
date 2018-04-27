(declare-sort Obj)
(declare-sort IRI)
(declare-sort Set)
(declare-datatypes
 ()
 ((Val (AnObj (obj Obj))
        (ANum (num Real))
        (AnInt (int Int))
        (ABool (bool Bool))
        (AString (str String))
        (AnIRI (iri IRI))
        (ASet (set Set))
        None)
  (IntExpr
   TheInt_e
   (Int_e (int_e Int))
   )
  (BoolIntExpr
   (And_e (and_a BoolIntExpr) (and_b BoolIntExpr))
   ;; (Or_e (or_e Expr Expr))
   (Gt_e (gt_a IntExpr) (gt_b IntExpr))
   (Lt_e (lt_a IntExpr) (lt_b IntExpr)))))

(declare-fun obj_get (Obj String) Val)

(define-fun is-Type ((x Val)) Bool
  (and
   (is-AnObj x)
   (is-ABool (obj_get (obj x) "canBeNone"))
   (is-ABool (obj_get (obj x) "canBeBool"))))

(define-fun app-int-expr ((it Int) (expr IntExpr)) Int
  (match expr
         ((TheInt_e it)
          ((Int_e x) x))))

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

