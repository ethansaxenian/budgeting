import axios from 'axios';
import { Category, Month, Plan, Transaction, TransactionType } from './types';
import { dateToStr } from './utils';

const api = axios.create({
  baseURL: 'http://127.0.0.1:8000/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getMonths = async (): Promise<Month[]> => {
  try {
    const response = await api.get('/months');
    return response.data;
  } catch (error) {
    console.error('Error fetching months:', error);
    return [];
  }
};

export const getTransactions = async (
  monthId: string | null,
  type: TransactionType | null
): Promise<Transaction[]> => {
  try {
    const response = await api.get(
      `/transactions?month_id=${monthId}&transaction_type=${
        type === null ? null : type.toLowerCase()
      }`
    );
    return response.data;
  } catch (error) {
    console.error('Error fetching transactions:', error);
    return [];
  }
};

export const postTransaction = async (
  date: Date,
  amount: number,
  description: string,
  category: Category,
  transactionType: TransactionType
) => {
  try {
    await api.post('/transactions', {
      date: dateToStr(date),
      amount,
      description,
      category,
      type: transactionType.toLowerCase(),
    });
  } catch (error) {
    console.error('Error adding transaction:', error);
  }
};

export const deleteTransaction = async (id: number) => {
  try {
    await api.delete(`/transactions/${id}`);
  } catch (error) {
    console.error('Error deleting transaction:', error);
  }
};

export default api;


export const getPlans = async (
  monthId: string | null,
  type: TransactionType | null
): Promise<Plan[]> => {
  try {
    const response = await api.get(
      `/plans?month_id=${monthId}&transaction_type=${
        type === null ? null : type.toLowerCase()
      }`
    );
    return response.data;
  } catch (error) {
    console.error('Error fetching plans:', error);
    return [];
  }
};
