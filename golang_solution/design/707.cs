public class MyLinkedList 
{
    class DNode 
    {
        public int val;
        public DNode l;
        public DNode r;
        public DNode(int v) 
        {
        this.val = v;
        }
    }
    
    
    int len = 0;
    DNode head; 
    DNode tail;
        
    /** Initialize your data structure here. */
    public MyLinkedList() 
    {
        head = new DNode(-1);
        tail= new DNode(-1);
        
        tail.l = head;
        head.r = tail;
    }
    
    /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
    public int Get(int index) 
    {
        if (index < 0 || index + 1 > len) 
            return -1;
        DNode p = head;
        while (index >= 0) 
        {
            p = p.r;
            index--;
        }
        
        return p.val;
    }
    
    /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
    public void AddAtHead(int val) 
    {
        DNode insertNode = new DNode(val);
        head.r.l = insertNode;
        insertNode.r = head.r;
        head.r = insertNode;
        insertNode.l = head;
        len++;
    }
    
    /** Append a node of value val to the last element of the linked list. */
    public void AddAtTail(int val) 
    {
        DNode insertNode = new DNode(val);
        tail.l.r = insertNode;
        insertNode.l = tail.l.r;
        tail.l = insertNode;
        insertNode.r = tail;
        len++;  
    }
    
    /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
    public void AddAtIndex(int index, int val) 
    {   
        if (index == len) {
            AddAtTail(val);
            return;
        } else if (index < 0) {
            AddAtHead(val);
            return;
        } else if (index > len) {
            return;
        }
        DNode aN = new DNode(val);
        DNode p = head;
        while (index > 0)
        {
            p = p.r;
            index--;
        }
        p.r.l = aN;
        aN.r = p.r;
        p.r = aN;
        aN.l = p;
        len++;
    }
    
    /** Delete the index-th node in the linked list, if the index is valid. */
    public void DeleteAtIndex(int index) 
    {   
        if (index + 1 > len)
            return;
        DNode p = head;
        while (index > 0)
        {
            p = p.r;
            index--;
        }
        
        DNode a = p;
        DNode b = p.r.r;
        
        a.r = b;
        b.l = a;
        len--;
    }
}



/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList obj = new MyLinkedList();
 * int param_1 = obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */