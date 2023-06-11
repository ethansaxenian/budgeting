import { FC } from 'react';
import TransactionsTable from './TransactionsTable';

type MonthPageProps = {
  monthId: string;
};

const MonthPage: FC<MonthPageProps> = ({ monthId }) => {
  return (
    <div>
      <h2>Month: {monthId}</h2>

      <h3>Income</h3>
      <TransactionsTable monthId={monthId} type="income" />

      <h3>Expense</h3>
      <TransactionsTable monthId={monthId} type="expense" />
    </div>
  );
};

export default MonthPage;
