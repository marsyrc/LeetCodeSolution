package golang_solution.trie;

import java.util.*;

class AutocompleteSystem {

    private Trie trie;

    public AutocompleteSystem(String[] sentences, int[] times) {
        trie = new Trie();
        int n = sentences.length;
        for (int i = 0; i < n; ++i)
            trie.insert(sentences[i], times[i]);
    }

    public List<String> input(char c) {
        return trie.searchChar(c);
    }
}

class Trie {
    private TrieNode root;
    private TrieNode current;
    private PriorityQueue<TrieNode> pq;
    private StringBuilder sb;

    public Trie() {
        root = new TrieNode();
        current = root;
        pq = new PriorityQueue<TrieNode>(4,
                (o1, o2) -> o1.times - o2.times != 0 ? o1.times - o2.times : o2.sentence.compareTo(o1.sentence));
        sb = new StringBuilder();
    }

    public void insert(String sentence, int times) {
        int n = sentence.length();
        for (int i = 0; i < n; i++) {
            insertNewChar(sentence.charAt(i));
        }
        current.sentence = sentence;
        current.times += times;
        current = root;
    }

    private boolean insertNewChar(char c) {
        int charNo = c == ' ' ? 26 : c - 'a';
        if (current.children[charNo] != null) {
            current = current.children[charNo];
            return false;
        }
        current.children[charNo] = new TrieNode();
        current = current.children[charNo];
        return true;
    }

    public List<String> searchChar(char c) {
        List<String> result = new LinkedList<>();
        if (c == '#') {
            current.sentence = sb.toString();
            ++current.times;
            current = root;
            sb.delete(0, sb.length());
            return result;
        }

        sb.append(c);

        if (insertNewChar(c))
        return result;

        searchSentence(current);
        while (pq.peek() != null)
            result.add(0, pq.poll().sentence);

        return result;
    }

    //dfs 搜索句子 加入pq
    public void searchSentence(TrieNode node) {
        if (node.sentence != null) {
            pq.offer(node);
            if (pq.size() > 3) {
                pq.poll();
            }
        }

        for (TrieNode child : node.children) {
            if (child != null) {
                searchSentence(child);
            }
        }
    }
}

class TrieNode {
    public TrieNode[] children;
    public String sentence;
    public int times;

    public TrieNode() {
        times = 0;
        sentence = null;
        children = new TrieNode[27];
    }
}
