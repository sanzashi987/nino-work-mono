export interface Node<T> {
  next: Node<T> | null;
  node: T;
  key: string;
}
class LinkList<T> {
  length = 0;

  head: Node<T> | null = null;

  append(node: Node<T>): void {
    if (this.head === null) {
      this.head = node;
    } else {
      let current = this.head;
      while (current.next) {
        current = current.next;
      }
      current.next = node;
    }
    this.length++;
  }

  * transverse(): any {
    let node = this.head;
    while (node) {
      yield node;
      node = node.next;
    }
  }

  search(key: string): Node<T> | null {
    let node = this.head;
    while (node !== null && node.key !== key) {
      node = node.next;
    }
    return node;
  }
}
export class LinkNode<T> implements Node<T> {
  node: T;

  key: string;

  next: Node<T> | null = null;

  constructor(node: T, key: string, next: Node<T> | null = null) {
    this.node = node;
    this.next = next;
    this.key = key;
  }
}

export default LinkList;
