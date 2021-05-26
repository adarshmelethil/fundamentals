#!/usr/local/bin/sbcl --script

; (loop for i upto 10 collect (write-line "Hello, World!"))
(defun meaning (life)
  "Return the computed meaning of LIFE"
  (let ((meh "abc"))
    ;; Invoke krakaboom
    (loop :for x :across meh
       :collect x)))


