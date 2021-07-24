# BudgetPass
budgetpass: a low-budget, likely insecure password manager written in Go

## Installation
Clone the repo, run ``go build bpass.go``, and then copy the resulting ``bpass`` binary into your ``$PATH``.

## Usage
### Basic Explanation
budgetpass stores each of your passwords in an encrypted file, where the name
of the file is the name you gave that password. The encryption uses AES, where
your master password is the key. So, unless somebody has your master password,
they won't be able to decrypt your password file, and hence will not have access
to any of your passwords.

The directory in which these password files are stored is
``$HOME/.local/share/bpass`` by default, but can be modified
via the ``$BP_HOME`` variable.

### Commands
- ``bpass new <password name>`` - create a new password. You will need quotation marks if your name contains spaces, but I would recommend against doing this, because filenames with spaces are ugly. Instead, use dashes, underscores, or camelCase.
- ``bpass get <password name>`` - retrieve a password. Again, spaces require quotes.
