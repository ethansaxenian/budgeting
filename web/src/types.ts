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

export const CategoryByType = {
  [TransactionType.Income]: [
    Category.Paycheck,
    Category.Interest,
    Category.Bonus,
    Category.Other,
  ],
  [TransactionType.Expense]: [
    Category.Food,
    Category.Gifts,
    Category.Medical,
    Category.Home,
    Category.Transportation,
    Category.Personal,
    Category.Utilities,
    Category.Savings,
    Category.Transportation,
    Category.Other,
  ],
};

export interface Month {
  id: string;
  month_id: string;
  starting_balance: number;
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

export interface Plan {
  id: number;
  amount: number;
  type: TransactionType;
  category: Category;
  month: number;
  year: number;
  monthId: number;
}
