import React, { useEffect, useState } from 'react';
import { getMonths } from './api';
import MonthPage from './MonthPage';
import {
  Container,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
} from '@chakra-ui/react';
import { Month } from './types';

const App: React.FC = () => {
  const [months, setMonths] = useState<Month[]>([]);
  const [selectedMonth, setSelectedMonth] = useState(0);

  useEffect(() => {
    const fetchMonths = async () => {
      try {
        const data = await getMonths();
        setMonths(data);
        const now = new Date();
        setSelectedMonth(
          data
            .map(({ id }) => id)
            .indexOf(`${now.getMonth() + 1}-${now.getFullYear()}`)
        );
      } catch (error) {
        console.error('Error fetching months:', error);
      }
    };

    fetchMonths();
  }, []);

  return (
    <Container minW="1000">
      <Tabs onChange={(index) => setSelectedMonth(index)} index={selectedMonth}>
        <TabList>
          {months.map((month) => (
            <Tab key={month.id}>{month.id}</Tab>
          ))}
        </TabList>
        <TabPanels>
          {months.map((month) => (
            <TabPanel key={month.id}>
              <MonthPage month={month} />
            </TabPanel>
          ))}
        </TabPanels>
      </Tabs>
    </Container>
  );
};

export default App;
