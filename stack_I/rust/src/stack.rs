pub struct Stack<V> {
    data: Vec<V>,
}

impl<V> Stack<V> {
    pub fn size(&self) -> usize {
        self.data.len()
    }

    pub fn is_empty(&self) -> bool {
        self.data.len() == 0
    }

    pub fn push(&mut self, value: V) {
        self.data.push(value);
    }

    pub fn pop(&mut self) {
        self.data.pop();
    }

    pub fn top(&mut self) -> Option<&V> {
        self.data.last()
    }
}

pub fn new<V>() -> Stack<V> {
    return Stack { data: vec![] };
}

#[cfg(test)]
mod tests {
    use crate::stack;

    #[derive(Debug)]
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

        let expected = email_income.top().expect("missing expected");

        assert_eq!(expected.from, "c");
        assert_eq!(expected.to, "d");
    }

    #[test]
    fn memory_usage() {
        let mut email_income = stack::new::<Email>();

        assert_eq!(std::mem::size_of_val(&*email_income.data), 0);

        for _n in 0..1_000_000 {
            email_income.push(Email { from: "a", to: "b" });
        }

        assert_eq!(std::mem::size_of_val(&*email_income.data), 32000000);

        for _n in 0..1_000_000 {
            email_income.pop();
        }

        assert_eq!(std::mem::size_of_val(&*email_income.data), 0);
    }
}
