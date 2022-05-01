<script>
    import Login from "$lib/Login/index.svelte";
    import { page } from "$app/stores";
    let isAuthenticated = false;

    /*
Copyright © 2019 W3C and Jeff Carpenter \<jeffcarp@chromium.org\>

Both the original source code and new contributions in this repository are released under the [3-Clause BSD license](https://opensource.org/licenses/BSD-3-Clause).

# The 3-Clause BSD License

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
        */
    function atob(data) {
        if (arguments.length === 0) {
            throw new TypeError("1 argument required, but only 0 present.");
        }

        // Web IDL requires DOMStrings to just be converted using ECMAScript
        // ToString, which in our case amounts to using a template literal.
        data = `${data}`;
        // "Remove all ASCII whitespace from data."
        data = data.replace(/[ \t\n\f\r]/g, "");
        // "If data's length divides by 4 leaving no remainder, then: if data ends
        // with one or two U+003D (=) code points, then remove them from data."
        if (data.length % 4 === 0) {
            data = data.replace(/==?$/, "");
        }
        // "If data's length divides by 4 leaving a remainder of 1, then return
        // failure."
        //
        // "If data contains a code point that is not one of
        //
        // U+002B (+)
        // U+002F (/)
        // ASCII alphanumeric
        //
        // then return failure."
        if (data.length % 4 === 1 || /[^+/0-9A-Za-z]/.test(data)) {
            return null;
        }
        // "Let output be an empty byte sequence."
        let output = "";
        // "Let buffer be an empty buffer that can have bits appended to it."
        //
        // We append bits via left-shift and or.  accumulatedBits is used to track
        // when we've gotten to 24 bits.
        let buffer = 0;
        let accumulatedBits = 0;
        // "Let position be a position variable for data, initially pointing at the
        // start of data."
        //
        // "While position does not point past the end of data:"
        for (let i = 0; i < data.length; i++) {
            // "Find the code point pointed to by position in the second column of
            // Table 1: The Base 64 Alphabet of RFC 4648. Let n be the number given in
            // the first cell of the same row.
            //
            // "Append to buffer the six bits corresponding to n, most significant bit
            // first."
            //
            // atobLookup() implements the table from RFC 4648.
            buffer <<= 6;
            buffer |= atobLookup(data[i]);
            accumulatedBits += 6;
            // "If buffer has accumulated 24 bits, interpret them as three 8-bit
            // big-endian numbers. Append three bytes with values equal to those
            // numbers to output, in the same order, and then empty buffer."
            if (accumulatedBits === 24) {
                output += String.fromCharCode((buffer & 0xff0000) >> 16);
                output += String.fromCharCode((buffer & 0xff00) >> 8);
                output += String.fromCharCode(buffer & 0xff);
                buffer = accumulatedBits = 0;
            }
            // "Advance position by 1."
        }
        // "If buffer is not empty, it contains either 12 or 18 bits. If it contains
        // 12 bits, then discard the last four and interpret the remaining eight as
        // an 8-bit big-endian number. If it contains 18 bits, then discard the last
        // two and interpret the remaining 16 as two 8-bit big-endian numbers. Append
        // the one or two bytes with values equal to those one or two numbers to
        // output, in the same order."
        if (accumulatedBits === 12) {
            buffer >>= 4;
            output += String.fromCharCode(buffer);
        } else if (accumulatedBits === 18) {
            buffer >>= 2;
            output += String.fromCharCode((buffer & 0xff00) >> 8);
            output += String.fromCharCode(buffer & 0xff);
        }
        // "Return output."
        return output;
    }
    /**
     * A lookup table for atob(), which converts an ASCII character to the
     * corresponding six-bit number.
     */

    const keystr =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

    function atobLookup(chr) {
        const index = keystr.indexOf(chr);
        // Throw exception if character is not in the lookup string; should not be hit in tests
        return index < 0 ? undefined : index;
    }
    let url = decodeURIComponent(atob($page.params["redirect"])) || "";
</script>

<div>
    <Login bind:isAuthenticated redirectURL={url} />
</div>
