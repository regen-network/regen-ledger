(require 'lsp-mode)
(require 'lsp-ui)

(define-derived-mode ceres-mode prog-mode "Ceres")

(lsp-define-stdio-client
 ;; This can be a symbol of your choosing. It will be used as a the
 ;; prefix for a dynamically generated function "-enable"; in this
 ;; case: lsp-prog-major-mode-enable
 ceres-mode
 "ceres"
 ;; This will be used to report a project's root directory to the LSP
 ;; server.
 (lambda () default-directory)
 ;; This is the command to start the LSP server. It may either be a
 ;; string containing the path of the command, or a list wherein the
 ;; car is a string containing the path of the command, and the cdr
 ;; are arguments to that command.
 '("ceres"))

(add-to-list 'auto-mode-alist '("\\.ceres\\'" . ceres-mode))

(add-hook 'ceres-mode #'lsp-ceres-enable)
(add-hook 'lsp-mode-hook 'lsp-ui-mode)
(add-hook 'ceres-mode-hook 'flycheck-mode)
