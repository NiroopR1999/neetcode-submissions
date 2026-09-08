class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        let encoded=''
     for (const str of strs) {
            // Append the length of the string followed by the delimiter and the string itself.
            encoded += `${str.length}#${str}`;
        }
        return encoded;
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        const decoded=[]
        let i=0
        while(i<str.length)
        {
            let delimeter=str.indexOf('#',i)
            let len=parseInt(str.substring(i,delimeter),10)
            i=delimeter+1
            decoded.push(str.substring(i,i+len))
            i+=len
            
        }
    return decoded
    }
}
