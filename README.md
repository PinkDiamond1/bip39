# bip39
Generate and validate BIP-39 mnemonic sentences

## disclaimer
>The use of this tool does not guarantee security or suitability
for any particular use. Please review the code and use at your own risk.

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

Download the code to a folder and cd to the folder, then run
```bash
go install
```
Install shell completion. For instance `bash` completion can be installed
by adding following line to your `.bashrc`:
```bash
source <(bip39 completion bash)
```

## generate new mnemonic
```bash
bip39 gen 
chaos mosquito citizen bone pencil crunch genuine this dice noise when digital grass urge jungle decade melody typical improve army couch degree anxiety rifle

bip39 gen --length=12
ankle already sight barely skate lazy admit estate plug sunset machine help

bip39 gen --length=15
royal guilt glad stereo match web muffin enough silent seek owner hungry expect resemble fault

bip39 gen --length=12 --language=1 # English
walk pudding hotel ordinary unknown detect simple typical leg ridge armor bitter

bip39 gen --length=12 --language=2 # Japanese
ただしい ほたる あける おおどおり からい はんい たいぐう めんきょ ふとん ねむたい てわけ もえる
```

## validate mnemonics
Entering correct mnemonic will produce no output indicating that the input is valid
```bash
bip39 validate
Enter mnemonic: mosquito citizen bone pencil crunch genuine this dice noise when digital grass urge jungle decade melody typical improve army couch degree anxiety rifle

bip39 validate --language=2
Enter mnemonic: ただしい ほたる あける おおどおり からい はんい たいぐう めんきょ ふとん ねむたい てわけ もえる
```

Entering invalid mnemonic will result in an error
```bash
bip39 validate 
Enter mnemonic: this is an invalid mnemonic
Error: invalid mnemonic
Usage:
  bip39 validate [flags]

Flags:
  -h, --help           help for validate
      --language int   Language (default 1)

Global Flags:
      --config string   config file (default is $HOME/.bip39.yaml)

Error: invalid mnemonic
```

Mnemonic is not any arbitrary list of words. The words come from a predefined list
and has a structure (last word has checksum). Altering the sequence of words
from a correct mnemonic will also result in error
```bash
bip39 gen --length=12
sausage unhappy suffer cost wedding air about maid future expand solar stumble

bip39 validate 
Enter mnemonic: unhappy sausage suffer cost wedding air about maid future expand solar stumble
Error: invalid mnemonic
Usage:
  bip39 validate [flags]

Flags:
  -h, --help           help for validate
      --language int   Language (default 1)

Global Flags:
      --config string   config file (default is $HOME/.bip39.yaml)

Error: invalid mnemonic
```

## translate mnemonics between languages
Several wallets only support english mnemonics. Mnemonics can be translated between
languages in the sense that the underlying entropy is preserved.

```bash
bip39 gen --language=2
ちかい ばいばい いずみ たおる おとなしい とくべつ もくてき たりきほんがん ふっかつ うける ざいりょう むかえ むすう けもの ちいき いがい きさく こうない げぼく うわさ そそぐ こんぽん にうけ はんい
```
English language equivalent mnemonic can be obtained from the above mnemonic:
```bash
bip39 translate --from-language=2 --to-language=1
Enter mnemonic: ちかい ばいばい いずみ たおる おとなしい とくべつ もくてき たりきほんがん ふっかつ うける ざいりょう むかえ むすう けもの ちいき いがい きさく こうない げぼく うわさ そそぐ こんぽん にうけ はんい
napkin school another mass caution pole universe mix stay become flock tray trophy electric myth already coyote essay egg book left first quit shoot
```
