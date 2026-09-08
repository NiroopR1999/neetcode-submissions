class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        let n=strs.length
        let map=new Map()

        for(let i=0;i<n;i++)
        {
               let temp = strs[i].split('').sort().join('');
              if(map.get(temp))
              {
                    let arr=map.get(temp)
                    arr.push(strs[i])
                     map.set(temp,arr)
              }
              else{
                map.set(temp,[strs[i]])
              }
        }

        return Array.from(map.values());

    }
}
