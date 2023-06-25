import axios from 'axios';
import { Month, Plan, Transaction, TransactionType } from './types';

const api = axios.create({
  baseURL: `${import.meta.env.VITE_API_URL}/api`,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getMonths = async (): Promise<Month[]> => {
  try {
    const response = await api.get('/months/');
    return response.data;
  } catch (error) {
    console.error('Error fetching months:', error);
    return [];
  }
};

export const putMonth = async (
  id: number,
  monthId: string,
  startingBalance: number
) => {
  try {
    await api.put(`/months/${id}`, {
      month_id: monthId,
      starting_balance: startingBalance,
    });
  } catch (error) {
    console.error('Error fetching months:', error);
  }
};

export const getTransactions = async (
  monthId: number | null,
  type: TransactionType | null
): Promise<Transaction[]> => {
  try {
    const response = await api.get(
      `/transactions/?month_id=${monthId}&transaction_type=${
        type === null ? null : type.toLowerCase()
      }`
    );
    return response.data;
  } catch (error) {
    console.error('Error fetching transactions:', error);
    return [];
  }
};

export const postTransaction = async (transaction: Omit<Transaction, 'id'>) => {
  try {
    await api.post('/transactions/', transaction);
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

export const putTransaction = async (transaction: Transaction) => {
  try {
    await api.put(`/transactions/${transaction.id}`, {
      type: transaction.type,
      amount: transaction.amount,
      description: transaction.description,
      category: transaction.category,
      date: transaction.date,
      month_id: transaction.month_id,
    });
  } catch (error) {
    console.error('Error updating transaction:', error);
  }
};

export const getPlans = async (
  monthId: number | null,
  type: TransactionType | null
): Promise<Plan[]> => {
  try {
    const response = await api.get(
      `/plans/?month_id=${monthId}&transaction_type=${
        type === null ? null : type.toLowerCase()
      }`
    );
    return response.data;
  } catch (error) {
    console.error('Error fetching plans:', error);
    return [];
  }
};

export const getPlan = async (id: number): Promise<Plan> => {
  try {
    const response = await api.get(`/plans/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching plan:', error);
    return {} as Plan;
  }
};

export const putPlan = async (plan: Plan) => {
  try {
    await api.put(`/plans/${plan.id}`, {
      type: plan.type,
      amount: plan.amount,
      category: plan.category,
      month_id: plan.month_id,
    });
  } catch (error) {
    console.error('Error updating plan:', error);
  }
};

export default api;
