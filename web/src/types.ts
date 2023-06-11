export enum TransactionType {
  Income = 'income',
  Expense = 'expense',
}

export enum Category {
  Food = 'Food',
  Gifts = 'Gifts',
  Medical = 'Medical',
  Home = 'Home',
  Transportation = 'Transportation',
  Personal = 'Personal',
  Savings = 'Savings',
  Utilities = 'Utilities',
  Travel = 'Travel',
  Other = 'Other',
  Paycheck = 'Paycheck',
  Bonus = 'Bonus',
  Interest = 'Interest',
}

export interface Month {
  id: string;
  startingBalance: number;
  name: number;
  year: number;
}

export interface Transaction {
  id: number;
  amount: number;
  type: TransactionType;
  description: string;
  date: string;
  category: Category;
}
