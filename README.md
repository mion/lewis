# Lewis

Lewis is a tiny dialect of Lisp that interoperates with Go (à la Clojure).

# Installation

Make sure you have Go installed, then clone the repo:

	$ git clone https://github.com/mion/lewis.git

Setup the workspace (Go environment variables):
	
	$ cd lewis
	$ source setup_workspace.sh

Build it:

	$ sh build.sh


Start the REPL:

	$ lewis

# Features

You can import and use Go packages as such:
	
	(import math/big)
	(define x (NewInt 123))
	(define y (NewInt 456))
	(define z (NewInt 0))
	(z Exp (x y nil))		-- From Go's "math/big": func (z *Int) Exp(x, y, m *Int) *Int
