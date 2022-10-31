pub struct Stack<V: Clone> {
    data: Vec<V>,
}

impl<V: Clone> Stack<V> {
    pub fn size(&self) -> usize {
        self.data.len()
    }

    pub fn is_empty(&self) -> bool {
        self.size() == 0
    }

    pub fn push(&mut self, value: V) {
        self.data.push(value);
    }

    pub fn pop(&mut self) {
        self.data.pop();
    }

    pub fn top(&mut self) -> &V {
        self.data.last().expect("impossible to top")
    }

    pub fn memory(&self) -> usize {
        std::mem::size_of_val(&*self.data.clone())
    }
}

pub fn new<V: Clone>() -> Stack<V> {
    return Stack { data: vec![] };
}

#[cfg(test)]
mod tests {
    use crate::stack;

    #[derive(Clone, Debug)]
    struct Email<'a> {
        from: &'a str,
        to: &'a str,
    }

    #[test]
    fn stack() {
        let mut email_income = stack::new::<Email>();

        assert_eq!(email_income.is_empty(), true);

        email_income.push(Email { from: "a", to: "b" });
        email_income.push(Email { from: "c", to: "d" });
        email_income.push(Email { from: "e", to: "f" });
        email_income.push(Email { from: "g", to: "h" });

        email_income.pop();
        email_income.pop();

        assert_eq!(email_income.size(), 2);

        let expected = email_income.top();

        assert_eq!(expected.from, "c");
        assert_eq!(expected.to, "d");
    }

    #[test]
    fn memory_usage() {
        let mut email_income = stack::new::<Email>();

        assert_eq!(email_income.memory(), 0);

        for _n in 0..1_000_000 {
            email_income.push(Email { from: "a", to: "b" });
        }

        assert_eq!(email_income.memory(), 32000000);

        for _n in 0..1_000_000 {
            email_income.pop();
        }

        assert_eq!(email_income.memory(), 0);
    }
}
