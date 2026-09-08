/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

class Solution {
    /**
     * @param {ListNode[]} lists
     * @return {ListNode}
     */
    mergeKLists(lists) {
        if (!lists || lists.length === 0) return null;

        const minHeap = new MinHeap((a, b) => a.val - b.val);

        for (let list of lists) {
            if (list) minHeap.insert(list);
        }

        const dummy = new ListNode(0);
        let current = dummy;

        while (!minHeap.isEmpty()) {
            const node = minHeap.extract();
            current.next = node;
            current = current.next;
            if (node.next) minHeap.insert(node.next);
        }

        return dummy.next;
    }
}

class MinHeap {
    constructor(compare) {
        this.heap = [];
        this.compare = compare;
    }

    insert(val) {
        this.heap.push(val);
        this.bubbleUp();
    }

    extract() {
        if (this.heap.length === 1) return this.heap.pop();
        const top = this.heap[0];
        this.heap[0] = this.heap.pop();
        this.bubbleDown();
        return top;
    }

    bubbleUp() {
        let index = this.heap.length - 1;
        while (index > 0) {
            const parentIndex = Math.floor((index - 1) / 2);
            if (this.compare(this.heap[index], this.heap[parentIndex]) >= 0) break;
            [this.heap[index], this.heap[parentIndex]] = [this.heap[parentIndex], this.heap[index]];
            index = parentIndex;
        }
    }

    bubbleDown() {
        let index = 0;
        const length = this.heap.length;
        while (true) {
            const leftChildIndex = 2 * index + 1;
            const rightChildIndex = 2 * index + 2;
            let smallest = index;

            if (leftChildIndex < length && this.compare(this.heap[leftChildIndex], this.heap[smallest]) < 0) {
                smallest = leftChildIndex;
            }

            if (rightChildIndex < length && this.compare(this.heap[rightChildIndex], this.heap[smallest]) < 0) {
                smallest = rightChildIndex;
            }

            if (smallest === index) break;
            [this.heap[index], this.heap[smallest]] = [this.heap[smallest], this.heap[index]];
            index = smallest;
        }
    }

    isEmpty() {
        return this.heap.length === 0;
    }
}
