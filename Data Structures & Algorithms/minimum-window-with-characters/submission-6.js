class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {string}
     */
    minWindow(s, t) {
        const x=s.length
        const y=t.length

        if(y===0 
        || y>x) return ''

        const required=new Map()
        const window=new Map()
        let met=0, left=0, right=0, min=Infinity, res=[-1,-1]

        for(const item of t)
        {
            required.set(item,(required.get(item)|| 0)+1)
        }

        const size=required.size

        while(right<x)
        {
              window.set(s[right],(window.get(s[right])|| 0)+1)

              if(required.has(s[right]) && window.get(s[right])===required.get(s[right]))
              {
                met++
              }
                while(met===size)
                {
                    if(right-left+1<min)
                    {
                        min=right-left+1
                        res=[left,right]
                    }

                    window.set(s[left],window.get(s[left])-1)

                    if(required.has(s[left]) && window.get(s[left])<required.get(s[left]))
                    {
                        met--
                    }
                    left++

                }
              right++
        }
        return res.includes(-1)? "" : s.slice(res[0],res[1]+1)
    }
}
