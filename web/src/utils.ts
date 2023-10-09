import { Category, Month, Plan, Transaction } from './types';

export const strToDate = (dateStr: string): Date => {
  const parts = dateStr.split('-');

  return new Date(
    parseInt(parts[0]),
    parseInt(parts[1]) - 1,
    parseInt(parts[2])
  );
};

export const dateToStr = (date: Date): string => {
  const month = `${date.getMonth() + 1}`.padStart(2, '0');
  const day = `${date.getDate()}`.padStart(2, '0');

  return `${date.getFullYear()}-${month}-${day}`;
};

export const formatAmount = (amount: number): string => {
  if (amount >= 0) {
    return `$${amount}`;
  }

  return `-$${Math.abs(amount)}`;
};

export const colorAmount = (amount: number): string => {
  return amount === 0 ? 'black' : amount > 0 ? 'green' : 'red';
};

export const round = (num: number): number => {
  return Math.round(num * 100) / 100;
};

export const isNumber = (num: string): boolean => {
  return !isNaN(num as unknown as number) && num !== '';
};

export const isDate = (date: string): boolean => {
  return /^\d{4}-(0[1-9]|1[012])-([0][1-9]|[12][0-9]|3[01])$/.test(date);
};

export const isNotEmpty = (val: string): boolean => {
  return !/^\s*$/.test(val);
};

export const sortTransactions = (
  transactions: Transaction[],
  field: keyof Transaction,
  asc: boolean
): Transaction[] => {
  switch (field) {
    case 'date':
      return transactions.sort(
        (a, b) =>
          strToDate((asc ? a : b).date).getDate() -
          strToDate((asc ? b : a).date).getDate()
      );

    case 'amount':
      return transactions.sort(
        (a, b) =>
          ((asc ? a : b).amount as number) - ((asc ? b : a).amount as number)
      );

    default:
      return transactions;
  }
};

export const sumPlan = (plan: Plan): number => {
  let sum = 0;
  for (const cat of Object.values(Category)) {
    sum += plan?.[cat] || 0;
  }
  return sum;
};

export const capitalize = (str: string): string => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};

export const compareMonths = (a: Month, b: Month): number => {
  const aParts = a.month_id.split('-');
  const bParts = b.month_id.split('-');

  const aDate = new Date(parseInt(aParts[1]), parseInt(aParts[0]));
  const bDate = new Date(parseInt(bParts[1]), parseInt(bParts[0]));

  return aDate.getTime() - bDate.getTime();
};
