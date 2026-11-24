# Bottom View of a Binary Tree

> Striver SDE Sheet 110 · Day 17 Binary Tree — tag nodes with a column number (left −1, right +1) and BFS; the last node seen per column is the visible one. Go, stdlib only, tested.

## Problem

Looking up at a binary tree from below, one node is visible per vertical column: the lowest. Columns are *horizontal distances* — root at 0, each left step −1, each right step +1 — so nodes from far-apart subtrees can share a column.

## Approach

Attach the column number to each node as it travels through a BFS queue. Because BFS visits strictly top-to-bottom, simply **overwriting** a per-column map entry on every visit leaves exactly the deepest (and, within a level, rightmost) node per column — the bottom view, with no depth comparisons stored anywhere. Track min/max column on the fly and read the map out left to right. The top view is the same walk with the overwrite flipped to "write only if absent". DFS can't substitute directly here: it visits deep-left before shallow-right, so "last wins" would lie.

**Complexity:** O(n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
