class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {

        const map= new Map()
        const n=strs.length

        for(let i=0;i<n;i++)
        {
            const temp=strs[i].split('').sort().join('')
            if(map.has(temp))
            {
                const arr=map.get(temp)
                arr.push(strs[i])
                map.set(temp,arr)
            }
            else{
                map.set(temp,[strs[i]])
            }
        }
        return Array.from(map.values())
    }
}
