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

## 🧠 Enterprise Implementation Notes

- We implemented a resizable array by manually handling memory via slice doubling.
- The code uses idiomatic Go patterns for modularity and error handling.
- CI-tested with GitHub Actions and documented for maintainability.
- Scalable architecture suitable for extending to advanced structures like Lists, HashMaps, etc.
