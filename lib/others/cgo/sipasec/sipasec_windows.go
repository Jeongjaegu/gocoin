// +build windows

package sipasec

/*
See these pages for some info on how to install MSYS2 and mingw64:
 * http://www.msys2.org/
 * https://stackoverflow.com/questions/30069830/how-to-install-mingw-w64-and-msys2#30071634

After having MSYS2 and Mingw64 installed, you will have to manually install (pacman -S <pkgname>)
the following packages:
 * make
 * autoconf
 * libtool
 * automake
 * ... probably something else, which I don't remember ATM

At this point you should be able to use "MSYS2 MinGW 64-bit" shell to make and install ("make
install"):
 * sipa's secp256k1 library - https://github.com/bitcoin/bitcoin/tree/master/src/secp256k1
 * libgmp - https://gmplib.org/

Having the two packages proparly installed withing your MSYS MinGW 64-bit envirionment, you shoule
be able to execute "go test" in this folder without any problems.

*/

/*
#cgo LDFLAGS: -lsecp256k1 -lgmp

#include <stdio.h>
#include <string.h>
#include "secp256k1.h"

static secp256k1_context *ctx;

static void secp256k1_start() {
	ctx = secp256k1_context_create(SECP256K1_CONTEXT_VERIFY);
}

static int ecdsa_signature_parse_der_lax(const secp256k1_context* ctx, secp256k1_ecdsa_signature* sig, const unsigned char *input, size_t inputlen) {
    size_t rpos, rlen, spos, slen;
    size_t pos = 0;
    size_t lenbyte;
    unsigned char tmpsig[64] = {0};
    int overflow = 0;

    secp256k1_ecdsa_signature_parse_compact(ctx, sig, tmpsig);

    if (pos == inputlen || input[pos] != 0x30) {
        return 0;
    }
    pos++;

    if (pos == inputlen) {
        return 0;
    }
    lenbyte = input[pos++];
    if (lenbyte & 0x80) {
        lenbyte -= 0x80;
        if (pos + lenbyte > inputlen) {
            return 0;
        }
        pos += lenbyte;
    }

    if (pos == inputlen || input[pos] != 0x02) {
        return 0;
    }
    pos++;

    if (pos == inputlen) {
        return 0;
    }
    lenbyte = input[pos++];
    if (lenbyte & 0x80) {
        lenbyte -= 0x80;
        if (pos + lenbyte > inputlen) {
            return 0;
        }
        while (lenbyte > 0 && input[pos] == 0) {
            pos++;
            lenbyte--;
        }
        if (lenbyte >= sizeof(size_t)) {
            return 0;
        }
        rlen = 0;
        while (lenbyte > 0) {
            rlen = (rlen << 8) + input[pos];
            pos++;
            lenbyte--;
        }
    } else {
        rlen = lenbyte;
    }
    if (rlen > inputlen - pos) {
        return 0;
    }
    rpos = pos;
    pos += rlen;

    if (pos == inputlen || input[pos] != 0x02) {
        return 0;
    }
    pos++;

    if (pos == inputlen) {
        return 0;
    }
    lenbyte = input[pos++];
    if (lenbyte & 0x80) {
        lenbyte -= 0x80;
        if (pos + lenbyte > inputlen) {
            return 0;
        }
        while (lenbyte > 0 && input[pos] == 0) {
            pos++;
            lenbyte--;
        }
        if (lenbyte >= sizeof(size_t)) {
            return 0;
        }
        slen = 0;
        while (lenbyte > 0) {
            slen = (slen << 8) + input[pos];
            pos++;
            lenbyte--;
        }
    } else {
        slen = lenbyte;
    }
    if (slen > inputlen - pos) {
        return 0;
    }
    spos = pos;
    pos += slen;

    while (rlen > 0 && input[rpos] == 0) {
        rlen--;
        rpos++;
    }
    if (rlen > 32) {
        overflow = 1;
    } else {
        memcpy(tmpsig + 32 - rlen, input + rpos, rlen);
    }

    while (slen > 0 && input[spos] == 0) {
        slen--;
        spos++;
    }
    if (slen > 32) {
        overflow = 1;
    } else {
        memcpy(tmpsig + 64 - slen, input + spos, slen);
    }

    if (!overflow) {
        overflow = !secp256k1_ecdsa_signature_parse_compact(ctx, sig, tmpsig);
    }
    if (overflow) {
        memset(tmpsig, 0, 64);
        secp256k1_ecdsa_signature_parse_compact(ctx, sig, tmpsig);
    }
    return 1;
}


static int secp256k1_verify(unsigned char *msg, unsigned char *sig, int siglen, unsigned char *pk, int pklen) {
	int result;
    secp256k1_pubkey pubkey;
	secp256k1_ecdsa_signature _sig;

	if (!secp256k1_ec_pubkey_parse(ctx, &pubkey, pk, pklen)) {
		return -1;
	}
	if (!ecdsa_signature_parse_der_lax(ctx, &_sig, sig, siglen)) {
		return -1;
	}

	secp256k1_ecdsa_signature_normalize(ctx, &_sig, &_sig);
	result = secp256k1_ecdsa_verify(ctx, &_sig, msg, &pubkey);

	return result;
}


*/
import "C"
import "unsafe"


// Verify ECDSA signature
func EC_Verify(pkey, sign, hash []byte) int {
	return int(C.secp256k1_verify((*C.uchar)(unsafe.Pointer(&hash[0])),
		(*C.uchar)(unsafe.Pointer(&sign[0])), C.int(len(sign)),
		(*C.uchar)(unsafe.Pointer(&pkey[0])), C.int(len(pkey))))
}

func init() {
	C.secp256k1_start()
}