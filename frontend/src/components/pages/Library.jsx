import React, { useState, useEffect } from 'react';
import { Box, SimpleGrid, Image, Text } from '@chakra-ui/react';
import { runtime } from '@wails/runtime';

// BookCard component to display each book
const BookCard = ({ book }) => (
  <Box borderWidth="1px" borderRadius="lg" overflow="hidden" p={4}>
    <Image src={book.CoverImagePath} alt={`Cover of ${book.Title}`} />
    <Box p="6">
      <Box display="flex" alignItems="baseline">
        <Text fontWeight="bold" textTransform="uppercase" fontSize="lg" lineHeight="tight" isTruncated>
          {book.Title}
        </Text>
      </Box>
      <Box>
        <Text mt="1" fontWeight="semibold" as="h4" lineHeight="tight" isTruncated>
          by {book.Author}
        </Text>
      </Box>
    </Box>
  </Box>
);

// Library component to fetch and display books
const Library = () => {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    // Use Wails runtime to call the GetBooks method in the Go backend
    runtime.Invoke("GetBooks").then(setBooks).catch(console.error);
  }, []);

  return (
    <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} spacing={10}>
      {books.map((book) => (
        <BookCard key={book.ID} book={book} />
      ))}
    </SimpleGrid>
  );
};

export default Library;
