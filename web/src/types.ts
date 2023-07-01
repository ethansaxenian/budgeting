export enum TransactionType {
  Income = 'income',
  Expense = 'expense',
}

export enum Category {
  Food = 'food',
  Gifts = 'gifts',
  Medical = 'medical',
  Home = 'home',
  Transportation = 'transportation',
  Personal = 'personal',
  Savings = 'savings',
  Utilities = 'utilities',
  Travel = 'travel',
  Other = 'other',
  Paycheck = 'paycheck',
  Bonus = 'bonus',
  Interest = 'interest',
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
    Category.Other,
  ],
};

export interface Month {
  id: number;
  month_id: string;
  starting_balance: number;
}

export interface Transaction {
  id: number;
  amount: number;
  type: TransactionType;
  description: string;
  date: string;
  category: Category;
  month_id: number;
}

export interface Plan {
  id: number;
  type: TransactionType;
  month_id: number;
  food: number;
  gifts: number;
  medical: number;
  home: number;
  transportation: number;
  personal: number;
  savings: number;
  utilities: number;
  travel: number;
  other: number;
  paycheck: number;
  bonus: number;
  interest: number;
}
