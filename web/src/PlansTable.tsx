import { FC, useEffect } from 'react';
import { getPlans, patchPlan } from './api';
import {
  HStack,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import {
  Category,
  CategoryByType,
  Plan,
  Transaction,
  TransactionType,
} from './types';
import {
  capitalize,
  colorAmount,
  formatAmount,
  isNumber,
  round,
} from './utils';
import EditableField from './EditableField';

interface PlansTableProps {
  monthId: number;
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
      const data = await getPlans(monthId, null);
      setPlans(data);
    };

    fetchPlans();
  }, [monthId, type, setPlans]);

  const updatePlanById = async (
    id: number,
    amount: string,
    category: Category
  ) => {
    await patchPlan(id, { [category]: parseFloat(amount) });
    const data = await getPlans(monthId, type);
    setPlans(data);
  };

  const [plan] = plans.filter((plan) => plan.type === type);

  return (
    <TableContainer>
      <Table fontSize="12px">
        <Thead>
          <Tr>
            <Th>Category</Th>
            <Th>Planned</Th>
            <Th>Actual</Th>
            <Th>Difference</Th>
          </Tr>
        </Thead>
        <Tbody>
          {Object.values(CategoryByType[type]).map((category) => {
            const actual = round(
              transactions.reduce(
                (sum, t) =>
                  sum +
                  (t.type === type && t.category === category ? t.amount : 0),
                0
              )
            );

            const diff = round((plan?.[category] || 0) - actual);

            return (
              <Tr key={`${category}-${plan?.id}`}>
                <Td w="30px">{capitalize(category)}</Td>
                <Td>
                  <HStack>
                    <Text m={-2} p={0}>
                      $
                    </Text>
                    <EditableField
                      initialValue={plan?.[category] || 0}
                      onSubmit={(val) => updatePlanById(plan.id, val, category)}
                      placeholder="00.00"
                      validate={isNumber}
                    />
                  </HStack>
                </Td>
                <Td w="30px">{formatAmount(actual)}</Td>
                <Td w="30px" color={colorAmount(diff)}>
                  {formatAmount(diff)}
                </Td>
              </Tr>
            );
          })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default PlansTable;
