# bip39
Generate and validate BIP-39 mnemonic sentences

## disclaimer
>The use of this tool does not guarantee security or suitability
for any particular use. Please review the code and use at your own risk.

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

```bash
go install github.com/kubetrail/bip39@latest
```
Install shell completion. For instance `bash` completion can be installed
by adding following line to your `.bashrc`:
```bash
source <(bip39 completion bash)
```

## generate new mnemonic
```bash
bip39 gen
```
```text
chaos mosquito citizen bone pencil crunch genuine this dice noise when digital grass urge jungle decade melody typical improve army couch degree anxiety rifle
```

```bash
bip39 gen --length=12
```
```text
ankle already sight barely skate lazy admit estate plug sunset machine help
```

```bash
bip39 gen --length=15
```
```text
royal guilt glad stereo match web muffin enough silent seek owner hungry expect resemble fault
```

```bash
bip39 gen --length=12 --language=English
```
```text
walk pudding hotel ordinary unknown detect simple typical leg ridge armor bitter
```

```bash
bip39 gen --length=12 --language=Japanese
```
```text
ただしい ほたる あける おおどおり からい はんい たいぐう めんきょ ふとん ねむたい てわけ もえる
```

## validate mnemonics
Entering correct mnemonic will produce no output indicating that the input is valid
```bash
bip39 validate
```
```text
Enter mnemonic: mosquito citizen bone pencil crunch genuine this dice noise when digital grass urge jungle decade melody typical improve army couch degree anxiety rifle
```

```bash
bip39 validate --language=Japanese
```
```text
Enter mnemonic: ただしい ほたる あける おおどおり からい はんい たいぐう めんきょ ふとん ねむたい てわけ もえる
```

Entering invalid mnemonic will result in an error
```bash
bip39 validate
```
```text
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
```
```text
sausage unhappy suffer cost wedding air about maid future expand solar stumble
```

```bash
bip39 validate 
```
```text
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
bip39 gen --language=Japanese
```
```text
ちかい ばいばい いずみ たおる おとなしい とくべつ もくてき たりきほんがん ふっかつ うける ざいりょう むかえ むすう けもの ちいき いがい きさく こうない げぼく うわさ そそぐ こんぽん にうけ はんい
```

English language equivalent mnemonic can be obtained from the above mnemonic:
```bash
bip39 translate --from-language=Japanese --to-language=English
```
```text
Enter mnemonic: ちかい ばいばい いずみ たおる おとなしい とくべつ もくてき たりきほんがん ふっかつ うける ざいりょう むかえ むすう けもの ちいき いがい きさく こうない げぼく うわさ そそぐ こんぽん にうけ はんい
napkin school another mass caution pole universe mix stay become flock tray trophy electric myth already coyote essay egg book left first quit shoot
```

## generate hex seed from mnemonics
Hex seeds are directly used for key generation. These hex seeds are generated
using mnemonic data and optional passphrase.

> All mnemonics are translated to English before seed is generated
> except when mnemonic validation is explicitly switched off

For instance, let's generate a Japanese language mnemonic and translate it to English
```bash
bip39 gen --language=japanese --length=12
```
```text
もくようび りきさく どうかん はつおん せんよう ぐんて げいじゅつ ぴっちり さとし こまる どぶがわ えつらん
```

```bash
bip39 translate --from-language=japanese --to-language=english
```
```text
Enter mnemonic: もくようび りきさく どうかん はつおん せんよう ぐんて げいじゅつ ぴっちり さとし こまる どぶがわ えつらん
unknown wasp pipe select knock divide doll smoke fringe feature present bright
```

We can see that the seed generated using both these mnemonics are the same

```bash
bip39 seed --language=english
```
```text
Enter mnemonic: unknown wasp pipe select knock divide doll smoke fringe feature present bright
29cc2a90e4f6cdd029cd1ff389c950ddcc1d895031c8cb3182af799ff37f69e9fe0f0de8b3a5d808fc9fe2e773c1a005d8d00c9a1a7f7d7a17e870234980e62e
```


```bash
bip39 seed --language=japanese
```
```text
Enter mnemonic: もくようび りきさく どうかん はつおん せんよう ぐんて げいじゅつ ぴっちり さとし こまる どぶがわ えつらん
29cc2a90e4f6cdd029cd1ff389c950ddcc1d895031c8cb3182af799ff37f69e9fe0f0de8b3a5d808fc9fe2e773c1a005d8d00c9a1a7f7d7a17e870234980e62e
```