import { useEffect, useState, FC } from 'react';
import api from './apiConfig';

interface Transaction {
  id: number;
  amount: number;
  type: string;
  description: string;
  date: string;
  category: string;
}

interface TransactionsTableProps {
  monthId: string;
  type: string;
}

const TransactionsTable: FC<TransactionsTableProps> = ({ monthId, type }) => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await api.get(
          `/transactions?month_id=${monthId}&transaction_type=${type}`
        );
        setTransactions(response.data);
      } catch (error) {
        console.error('Error fetching transactions:', error);
      }
    };

    fetchTransactions();
  }, [monthId, type]);

  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>Date</th>
            <th>Amount</th>
            <th>Type</th>
            <th>Description</th>
            <th>Category</th>
          </tr>
        </thead>
        <tbody>
          {transactions.map((transaction) => (
            <tr key={transaction.id}>
              <td>{transaction.date}</td>
              <td>{transaction.amount}</td>
              <td>{transaction.type}</td>
              <td>{transaction.description}</td>
              <td>{transaction.category}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TransactionsTable;
