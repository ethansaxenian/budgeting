import { FC, useEffect } from 'react';
import { getPlans } from './api';
import {
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import { CategoryByType, Plan, Transaction, TransactionType } from './types';
import { colorAmount, formatAmount, round } from './utils';

interface PlansTableProps {
  monthId: string;
  type: TransactionType;
  plans: Plan[];
  transactions: Transaction[];
  setPlans: (plans: Plan[]) => void;
}

const PlansTable: FC<PlansTableProps> = ({
  monthId,
  type,
  plans,
  transactions,
  setPlans,
}) => {
  useEffect(() => {
    const fetchPlans = async () => {
      const data = await getPlans(monthId, type);
      setPlans(data);
    };

    fetchPlans();
  }, [monthId, type, setPlans]);

  const planMap: Record<string, Record<string, number>> = {};

  for (const cat of CategoryByType[type]) {
    const plansForCategory = plans.filter((p) => p.category === cat);
    const sumPlansForCategory = round(
      plansForCategory.reduce((sum, p) => sum + p.amount, 0)
    );
    const transactionsForCategory = transactions.filter(
      (t) => t.category === cat
    );
    const sumTransactionsForCategory = round(
      transactionsForCategory.reduce((sum, t) => sum + t.amount, 0)
    );

    planMap[cat] = {
      planned: sumPlansForCategory,
      actual: sumTransactionsForCategory,
      diff: round(
        type === TransactionType.Expense
          ? sumPlansForCategory - sumTransactionsForCategory
          : sumTransactionsForCategory - sumPlansForCategory
      ),
    };
  }

  return (
    <TableContainer>
      <Table fontSize="12px">
        <Thead>
          <Tr>
            <Th>Category</Th>
            <Th>Planned</Th>
            <Th>Actual</Th>
            <Th>Difference</Th>
            <Th />
          </Tr>
        </Thead>
        <Tbody>
          {Object.entries(planMap).map(
            ([category, { planned, actual, diff }]) => (
              <Tr key={category}>
                <Td w="30px">{category}</Td>
                <Td w="30px">{formatAmount(planned)}</Td>
                <Td w="30px">{formatAmount(actual)}</Td>
                <Td w="30px" color={colorAmount(diff)}>
                  {formatAmount(diff)}
                </Td>
              </Tr>
            )
          )}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default PlansTable;
