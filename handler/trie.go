package handler

import (
  "strings"
)

/* 
 * 定义 Trie 树的节点
 * 使用 rune 支持 Unicode（中文、英文等）
 * 是否为一个敏感词的结尾
 */
type TrieNode struct {
  children map[rune]*TrieNode
  isEnd    bool
}

/* 创建一个新的 Trie 节点*/
func NewTrieNode() *TrieNode {
  return &TrieNode{
    children: make(map[rune]*TrieNode),
    isEnd:    false,
  }
}

/* 定义整个Trie树*/
type Trie struct {
  root *TrieNode
}

/* 创建一个新的 Trie 树*/
func NewTrie() *Trie {
  return &Trie{
    root: NewTrieNode(),
  }
}

/* 向 Trie 树中插入一个敏感词*/
func (t *Trie) Insert(word string) {
  node := t.root
  for _, ch := range word {
    if _, ok := node.children[ch]; !ok {
      node.children[ch] = NewTrieNode()
    }
    node = node.children[ch]
  }
  node.isEnd = true
}

/* 从 Trie 树中移除一个敏感词*/
func (t *Trie) Remove(word string) {
  if word == "" {
    return
  }
  runes := []rune(word)
  path := make([]struct {
    node *TrieNode
    ch   rune
  }, 0, len(runes))
  node := t.root
  for _, ch := range runes {
    child, ok := node.children[ch]
    if !ok {
      return
    }
    path = append(path, struct {
      node *TrieNode
      ch   rune
    }{node, ch})
    node = child
  }
  if !node.isEnd {
    return
  }
  node.isEnd = false
  for i := len(path) - 1; i >= 0; i-- {
    parent := path[i].node
    ch := path[i].ch
    child := parent.children[ch]
    if !child.isEnd && len(child.children) == 0 {
      delete(parent.children, ch)
    } else {
      break
    }
  }
}

/*敏感词过滤函数：检测文本中的敏感词并替换为 */
func (t *Trie) Filter(text string) string {
  var result strings.Builder
  runes := []rune(text)
  n := len(runes)

  for i := 0; i < n; {
    node := t.root
    j := i
    /*最后一个匹配成功的结束位置*/
    lastMatchEnd := -1

    /*从当前位置开始，逐个字符在 Trie 树中查找*/
    for j < n {
      ch := runes[j]
      if child, ok := node.children[ch]; ok {
        node = child
        if node.isEnd {
          lastMatchEnd = j // 找到一个敏感词结尾
        }
        j++
      } else {
        break
      }
    }

    if lastMatchEnd != -1 {
      /*替换*/
      result.WriteString("***") 
      i = lastMatchEnd + 1
    } else {
      /*没有敏感词，保留当前字符*/
      result.WriteRune(runes[i])
      i++
    }
  }
  return result.String()
}

/*敏感词过滤函数：检测文本中的敏感词并替换为 */
func (t *Trie) Check(text string) []string {
  var result []string
  runes := []rune(text)
  n := len(runes)

  for i := 0; i < n; {
    node := t.root
    j := i
    /*最后一个匹配成功的结束位置*/
    lastMatchEnd := -1

    /*从当前位置开始，逐个字符在 Trie 树中查找*/
    for j < n {
      ch := runes[j]
      if child, ok := node.children[ch]; ok {
        node = child
        if node.isEnd {
          lastMatchEnd = j // 找到一个敏感词结尾
        }
        j++
      } else {
        break
      }
    }

    if lastMatchEnd != -1 {
      result = append(result, string(runes[i : lastMatchEnd+1]))
      i = lastMatchEnd + 1
    } else {
      i++
    }
  }
  return result
}
