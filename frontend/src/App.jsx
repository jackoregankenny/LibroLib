import React from 'react';
import {
  ChakraProvider,
  extendTheme,
  Box,
  Flex,
  VStack,
  IconButton,
  Drawer,
  DrawerContent,
  DrawerOverlay,
  useDisclosure,
  DrawerCloseButton,
  Text,
  Icon,
  Link,
} from '@chakra-ui/react';
import { HamburgerIcon, InfoIcon, SearchIcon, AddIcon } from '@chakra-ui/icons';
import { BrowserRouter as Router, Routes, Route, Link as RouterLink } from 'react-router-dom';
import Home from './components/pages/Home';
import Library from './components/pages/Library';
import BookDetails from './components/pages/BookDetails';

// Custom theme for E Ink appearance
const theme = extendTheme({
  colors: {
    black: '#000000',
    white: '#FFFFFF',
    grey: '#CCCCCC',
  },
  components: {
    // Add component customizations if needed
  },
});

function SidebarItem({ icon, label, to }) {
  const { onClose } = useDisclosure();
  return (
    <Link as={RouterLink} to={to} onClick={onClose} style={{ textDecoration: 'none' }}>
      <Flex align="center" p="2">
        <Icon as={icon} mr="2" />
        <Text display={{ base: 'none', md: 'block' }}>{label}</Text>
      </Flex>
    </Link>
  );
}

function App() {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <ChakraProvider theme={theme}>
      <Router>
        <Box minH="100vh" display={{ md: 'flex' }}>
          {/* Drawer for mobile view */}
          <Drawer isOpen={isOpen} placement="left" onClose={onClose}>
            <DrawerOverlay />
            <DrawerContent>
              <DrawerCloseButton />
              <VStack spacing={4} m={4}>
                <SidebarItem to="/" icon={InfoIcon} label="Home" />
                <SidebarItem to="/library" icon={SearchIcon} label="Library" />
                <SidebarItem to="/add" icon={AddIcon} label="Add Book" />
                {/* Add more SidebarItem components as needed */}
              </VStack>
            </DrawerContent>
          </Drawer>

          {/* Sidebar for larger screens */}
          <Box
            as="nav"
            width={{ base: '0', md: '200px' }}
            bg="grey    "
            color="white"
            p={4}
            display={{ base: 'none', md: 'block' }}
          >
            <VStack spacing={4} align="stretch">
              <SidebarItem to="/" icon={InfoIcon} label="Home" />
              <SidebarItem to="/library" icon={SearchIcon} label="Library" />
              <SidebarItem to="/add" icon={AddIcon} label="Add Book" />
              {/* Add more SidebarItem components as needed */}
            </VStack>
          </Box>

          {/* Main Content */}
          <Box as="main" flex="1" p={4} bg="white" color="black">
            <IconButton
              icon={<HamburgerIcon />}
              aria-label="Open Menu"
              onClick={onOpen}
              size="lg"
              position="fixed"
              zIndex="overlay"
              m={2}
              display={{ md: 'none' }}
              colorScheme="whiteAlpha"
            />
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/library" element={<Library />} />
              <Route path="/add" element={<BookDetails />} />
              {/* Add more routes as needed */}
            </Routes>
          </Box>
        </Box>
      </Router>
    </ChakraProvider>
  );
}

export default App;
