class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        let encoded=''
        for(const str of strs)
        {
            encoded+=`${str.length}#${str}`
        }
        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        const decoded=[]
        let i=0
        const n=str.length
        while(i<n)
        {
            const delimeter=str.indexOf('#',i)
            const len=Number.parseInt(str.substring(i,delimeter),10)
            i=delimeter+1
            decoded.push(str.substring(i,i+len))
            i+=len
        }
        return decoded
    }
}
