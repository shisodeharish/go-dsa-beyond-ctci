# Chapter 25 – Dynamic Arrays

This chapter focuses on building a custom dynamic array data structure using fixed-size arrays, mimicking what languages like Python or JavaScript offer natively.

---

## 📌 Problem 25.1 – Implement Dynamic Array

Implement a dynamic array from scratch using only fixed-size arrays. Avoid using built-in `append()` methods.

### ✅ Supported Operations

- `append(x)`: Adds element `x` to the end of the array.
- `get(i)`: Returns the element at index `i`.
- `set(i, x)`: Updates the element at index `i` to value `x`.
- `size()`: Returns the number of elements in the array.
- `pop_back()`: Removes the last element in the array.

---

### 📘 Example 1 – Basic Append and Get
```python
d = DynamicArray()
d.append(1)
d.append(2)
d.get(0)  # returns 1
d.get(1)  # returns 2
d.size()  # returns 2
```

---

### 📘 Example 2 – Set
```python
d = DynamicArray()
d.append(1)
d.set(0, 10)
d.get(0)  # returns 10
```

---

### 📘 Example 3 – Pop Back
```python
d = DynamicArray()
d.append(1)
d.append(2)
d.pop_back()
d.size()   # returns 1
d.get(0)   # returns 1
```

---

### 🔒 Constraints

- All operations should support arrays with up to 10^6 elements.
- All integer values are between -10^9 and 10^9.

---

### 🧠 Solution Explanation (Golang with Visuals)

The core challenge of implementing a **dynamic array using only fixed-size arrays** is managing memory manually.

---

### 🚀 Key Principles

1. **Fixed-size internal storage**: We begin with a small capacity (e.g., 10).
2. **Dynamic resizing**: When the array is full, we:
   - Allocate a new array with double the capacity.
   - Copy all elements into the new array.
3. **Index safety**: Every index access is bounded by the current size (not capacity).

---

### 📈 Why Resizing Works Efficiently

When we double capacity on each overflow, the amortized cost of `append` is O(1):

```
Total cost of n appends:
  = n (copy during last resize)
  + n/2 (previous resize)
  + n/4 ...
  < 2n → O(n)

So, each append ≈ O(1) on average.
```

---

### 🧮 Shrinking Strategy for `pop_back()`

To save memory, we shrink the capacity when the array usage drops below 25%:

| Strategy         | Outcome                            |
|------------------|-------------------------------------|
| Never shrink     | ❌ Wastes memory                    |
| Shrink at 50%    | ❌ Can lead to immediate resizing   |
| ✅ Shrink at 25% | ✅ Balanced resizing, prevents bouncing |

---

### 🧊 Space Complexity

- At most **O(n)**, where `n` is the number of stored elements.
- After resize, worst-case memory usage is ~4n, still O(n).

---

### 🖼️ Illustration: Append Operation with Resizing

```plaintext
Before Append:
[1, 2, 3, _, _, _, _, _, _, _] (capacity = 10, size = 3)

Append(4)

After Append:
[1, 2, 3, 4, _, _, _, _, _, _] (size = 4)
```

If capacity is reached:

```plaintext
Before Append (Full):
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10] (capacity = 10)

→ Resize → New capacity = 20

After Resize:
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, _, _, _, _, _, _, _, _, _, _]
```

---

### 📌 Final Notes

- No `append()` used internally — resizing logic is handled manually.
- Capacity and size are clearly tracked.
- Safe, idiomatic Go implementation for practicing system-level memory control.

## 📌 Problem 25.2 – Extra Dynamic Array Operations

Enhance the previous `DynamicArray` by adding the following methods:

### ✅ Additional Operations

1. `pop(i)` – Removes the element at index `i`, shifts all elements after it to the left.
2. `contains(x)` – Returns `True` if element `x` exists in the array.
3. `insert(i, x)` – Inserts `x` at index `i`, shifting other elements right.
4. `remove(x)` – Removes first occurrence of `x`, shifts remaining elements, returns index or `-1` if not found.

---

### 📘 Example 1 – `pop(i)`
```python
d = DynamicArrayExtras()
d.append(1)
d.append(2)
d.append(3)
d.pop(1)     # returns 2
d.get(1)     # returns 3
d.size()     # returns 2
```

---

### 📘 Example 2 – `contains(x)`
```python
d = DynamicArrayExtras()
d.append(1)
d.append(2)
d.contains(1)  # returns True
d.contains(3)  # returns False
```

---

### 📘 Example 3 – `insert(i, x)`
```python
d = DynamicArrayExtras()
d.append(1)
d.append(2)
d.insert(1, 3)   # array becomes [1, 3, 2]
d.get(1)         # returns 3
```

---

### 📘 Example 4 – `remove(x)`
```python
d = DynamicArrayExtras()
d.append(1)
d.append(2)
d.append(2)
d.remove(2)      # returns 1
d.get(1)         # returns 2
```

---

### ⛓️ Constraints

- All operations should support arrays with up to 10^6 elements.
- All integer values are between -10^9 and 10^9.
---

### 🧠 Solution Explanation (Golang with Visuals)

Implementing extra operations on a dynamic array introduces **element shifting** and **sequential scans**. Here's how each is tackled idiomatically in Go.

---

### 1️⃣ `pop(i)` — Remove Element at Index

There are two scenarios:
- If `i == size-1`: We use `pop_back()` directly.
- Else: We **shift all elements left** after index `i`.

```plaintext
Before pop(2):
[1, 2, 3, 4, _]

After shifting:
[1, 2, 4, _, _] (size--)

Worst-case: O(n), especially when popping from index 0.
```

> ⚠️ Cannot amortize this: doing n pops from the front = O(n²)

---

### 2️⃣ `contains(x)` — Linear Scan

We scan the array until we find `x`.

```plaintext
for i := 0; i < size; i++ {
    if data[i] == x {
        return true
    }
}
return false
```

- Worst-case: O(n)
- ❗ Consider using hash sets or sorted arrays + binary search if faster lookup is needed.

---

### 3️⃣ `insert(i, x)` — Insert at Index

Insert requires **right-shifting** all elements from index `i` onward.

Steps:
1. Resize if at capacity.
2. Shift all elements right from index `i`.
3. Place `x` at `i`.

```plaintext
Before insert(1, 9):
[1, 2, 3, _, _]

Shift → [1, _, 2, 3, _]
Insert → [1, 9, 2, 3, _]
```

- Worst-case: O(n) when inserting at index 0.
- ❌ Not amortizable → n front inserts = O(n²)

---

### 4️⃣ `remove(x)` — Remove First Occurrence

This combines `contains()` and `pop(i)`.

Steps:
1. Scan to find index `i` where `x` exists.
2. Use `pop(i)` to shift elements left.
3. Return `i` or `-1` if not found.

```plaintext
Before: [1, 2, 3, 2, 4]
remove(2) → [1, 3, 2, 4, _] → returns 1
```

- Time: O(n) scan + O(n) shift → O(n)

---

### 🧠 Optimization Insights

| Operation     | Time Complexity | Amortized? | Comments |
|---------------|------------------|------------|----------|
| `append(x)`   | O(1)             | ✅         | Uses doubling |
| `pop_back()`  | O(1)             | ✅         | Shrinks at 25% |
| `pop(i)`      | O(n)             | ❌         | Shifting |
| `insert(i,x)` | O(n)             | ❌         | Right-shifting |
| `remove(x)`   | O(n)             | ❌         | Linear search + shift |
| `contains(x)` | O(n)             | ❌         | Linear search |

---

### 📦 Consider Alternatives

If frequent operations are needed at both ends or middle, consider:

- `deque` (double-ended queue) for O(1) front/back operations.
- `linked list` for constant-time inserts/removes.
- `hash table` or `set` for faster membership testing.

---

### 🏁 Final Notes

- Implementing these ops manually builds a solid foundation in memory handling.
- Efficient data layout, safe index management, and error handling make this solution idiomatic Go.
- Combined with GitHub Actions CI and clean architecture, this DSA module is robust and production-grade.

## 🧠 Enterprise Implementation Notes

- We implemented a resizable array by manually handling memory via slice doubling.
- The code uses idiomatic Go patterns for modularity and error handling.
- CI-tested with GitHub Actions and documented for maintainability.
- Scalable architecture suitable for extending to advanced structures like Lists, HashMaps, etc.
