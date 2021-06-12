public class Node
{
	public int val;
	public Node next;

	public Node() { }

	public Node(int _val)
	{
		val = _val;
		next = null;
	}

	public Node(int _val, Node _next)
	{
		val = _val;
		next = _next;
	}
}

public class Solution
{
	public Node Insert(Node head, int insertVal)
	{
		Node insert = new Node(insertVal);
		if (head == null)
		{
			insert.next = insert;
			return insert;
		}
		Node pre = head;
		Node cur = head.next;
		while (cur != head)
		{
			if (pre.val <= insertVal && insertVal <= cur.val)
			{
				break;
			}
			else if (pre.val > cur.val && insertVal <= cur.val)
			{
				break;
			}
			else if (pre.val > cur.val && insertVal >= pre.val)
			{
				break;
			}
			else
			{
				pre = pre.next;
				cur = cur.next;
			}
		}
		pre.next = insert;
		insert.next = cur;
		return head;
	}
}