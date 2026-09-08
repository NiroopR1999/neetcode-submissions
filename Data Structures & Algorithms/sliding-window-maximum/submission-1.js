class ListNode {
    constructor(value = null) {
        this.value = value;
        this.next = null;
        this.prev = null;
    }
}

class Deque {
    constructor() {
        this.head = new ListNode(); // Dummy head
        this.tail = new ListNode(); // Dummy tail
        this.head.next = this.tail;
        this.tail.prev = this.head;
    }

    push(val) {
        const node = new ListNode(val);
        node.prev = this.tail.prev;
        node.next = this.tail;
        this.tail.prev.next = node;
        this.tail.prev = node;
    }

    pop() {
        if (this.isEmpty()) return null;
        const node = this.tail.prev;
        node.prev.next = this.tail;
        this.tail.prev = node.prev;
        return node.value;
    }

    shift() {
        if (this.isEmpty()) return null;
        const node = this.head.next;
        node.next.prev = this.head;
        this.head.next = node.next;
        return node.value;
    }

    front() {
        if (this.isEmpty()) return null;
        return this.head.next.value;
    }

    back() {
        if (this.isEmpty()) return null;
        return this.tail.prev.value;
    }

    isEmpty() {
        return this.head.next === this.tail;
    }
}
class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    
    maxSlidingWindow(nums, k) {
            let n=nums.length
            if(n===0) return []
            let deque= new Deque()
            let res=[]

            for(let i=0;i<n;i++)
            {
                if(!deque.isEmpty() && i>=k && deque.front()===nums[i-k])
                {
                    deque.shift()
                }

                while(!deque.isEmpty() && nums[i]>deque.back())
                {
                    deque.pop()
                }

                deque.push(nums[i])
                if(i>=k-1)
                {
                    res.push(deque.front())
                }
            }
            return res
    }
}
