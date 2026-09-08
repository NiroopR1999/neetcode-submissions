class MinStack {
    constructor() {
        this.items=[]
    }

    /**
     * @param {number} val
     * @return {void}
     */
    push(val) {
        this.items.push(val)
    }

    /**
     * @return {void}
     */
    pop() {
        if(this.items.length===0) return 'underflow'
        this.items.pop()
    }

    /**
     * @return {number}
     */
    top() {
         if(this.items.length===0) return 'underflow'
        return this.items[this.items.length-1]
    }

    /**
     * @return {number}
     */
    getMin() {
        let min=Infinity

        for(const item of this.items)
        {
            if(item < min) min=item
        }
        return min
    }
}
