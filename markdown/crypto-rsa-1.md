The key (no pun intended) to understanding most cryptographic applications is a basic understanding of Asymetrical cryptograpy. Even in symmetrical key cryptography, the master key is usually passed through Asymmetrical mechanisms.

What is asymmetrical and symmetrical? Quick recap from your crypto course in uni...

<br>
### Symmetrical Schemes
<br>

Symmetrical refers to the process of decrypting a message using the reverse process of the one used to encrypt the message; meaning that the two processes are symmetrical.

The simplest example I can think:

Encrypt by multiplying by 2. Decrypt by dividing by 2.
<br>
<br>
`Message: 2 4 3 6`
<br>
`Encrypted: 4 8 6 12`
<br>
`Decrypted: 4/2 8/2 6/2 12/2`

In this scheme both receiver and sender must know the encryption scheme and key inorder for decryption to be possible. In this case the "scheme" is to multiply to encrypt, and divide to decrypt. Both of which use the *key* of value 2. If everyone party involved knows the scheme and the key, everyone can exchange messages. If an attacker knows the scheme, he must figure out the key.

Key question to ask: How do we pass around the key without anyone knowing it?
<br>
Let us compare to an Asymmetrical Scheme.

<br>
### Asymmetrical Schemes
<br>

The premise to symmetrical schemes is that both parties understand the cypher. Asymmetrical operates on the premise that only one party can decrypt but everyone can encrypt.

Addressing the question from the previous section, how do we pass around the key in an asymmetrical system? if everyone can encrypt, there is no need to pass a decryption key around. This is one of the greatest benefits to this scheme. The decryption key is know as the **Private Key**. 

Okay, if there is a *private key* to decrypt, what do we use to encrypt? Without getting into the math, we need to get a value that has a mathematical relation to this private key. Two numbers are generated, a **Public Key** and a so called **Modulus**.

Encryption looks like this:
<br>
`function of PubKey + ModKey(MessageToEncrypt) = Encrypted-Messages`

And decryption looks like this:
<br>
`function of PrivKey + ModKey(Encrypted-Message) = MessageToEncrypt`

The **ONLY** thing you must know from a high level is that it is **IMPOSSIBLE** by computational means to get the private key from knowing about the public key. With this assumption, everyone can know the Modulus and Public key and can never ever ever get to the private key by computation (this doesn't protect you against general carelessness) :P

<br>
### Lets recap
<br>

Symmetrical: One key known as the *private key*
Asymmetrical: Three keys, mathematically related to each other.

1. Private Key can not be cracked without getting to the physical key (ie. by breaking onto the machine where the key is located... but you probably have other problems at that point).
2. The three keys are not random, they are related to each other mathematically. The generation of these keys will be covered in a later post.

Some questions to ask yourself:

1. How do we trust a public key? We are encrypting data to be sent to the some person so how do we confirm this is the person we want to send this data to... especially if it is your bank account number?

<br>
### Conclussion
<br>

So, thats the bare, bare basics. How are asymmetrical and symmetrical keys related to each other in the real world. 

In asymmetrical encryption, the encryption can be given to anyone. It doesn't matter who gets this key. If we are communicating over an open network (aka. a network where anyone can listen to the conversation) we can send this public key without any fear.

HTTPS and SSH, for example, use Asymmetrical encryption to encrypt a key that is eventuall used for Symmetrical encryption. Symmetrical encryption/decryption is generally faster and much more practical in real-time scenarios.

Till next time,

Darius
