"use client";

import { useState, useEffect, ChangeEvent } from "react";
import { motion } from "motion/react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Command,
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
} from "@/components/ui/command";
import axios from "axios";
import { ExchangeRates, Token } from "@/types/types";
import Image from "next/image";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ClipLoader } from "react-spinners";
import { toast, ToastContainer } from "react-toastify";
import { ChevronsUpDown, Check } from "lucide-react";

const SWAP_API = "https://interview.switcheo.com/prices.json";

export default function SwapForm() {
  const [tokens, setTokens] = useState<Token[]>([]); // Array of deduplicated token objects
  const [prices, setPrices] = useState<ExchangeRates>({}); // Mapping: { [currency]: price }
  const [fromToken, setFromToken] = useState<string>("");
  const [toToken, setToToken] = useState<string>("");
  const [amount, setAmount] = useState<number | null>(null);
  const [convertedAmount, setConvertedAmount] = useState("0.00");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [showWalletOption, setShowWalletOption] = useState<boolean>(false);

  // Popover open states for comboboxes
  const [openFrom, setOpenFrom] = useState(false);
  const [openTo, setOpenTo] = useState(false);

  // Local search query states for each combobox
  const [fromSearchQuery, setFromSearchQuery] = useState("");
  const [toSearchQuery, setToSearchQuery] = useState("");

  useEffect(() => {
    setIsLoading(true);
    setError(null);
    axios
      .get(SWAP_API)
      .then(({ data }) => {
        // 1. Deduplicate tokens by currency using the latest date
        const dedupedTokens = data.reduce(
          (acc: Record<string, Token>, token: Token) => {
            if (
              !acc[token.currency] ||
              new Date(token.date) > new Date(acc[token.currency].date)
            ) {
              acc[token.currency] = token;
            }
            return acc;
          },
          {}
        );

        const exchangeRates = Object.fromEntries(
          Object.entries(dedupedTokens).map(([currency, token]) => [
            currency,
            (token as Token).price,
          ])
        );

        setPrices(exchangeRates);
        const tokensArr = Object.values(dedupedTokens) as Token[];
        setTokens(tokensArr);
      })
      .catch((error) => {
        setError(error.message);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  const handleSwap = () => {
    setFromToken(toToken);
    setToToken(fromToken);
  };

  const handleFromTokenChange = (value: string): void => {
    if (value === toToken) {
      setToToken(fromToken);
    }
    setFromToken(value);
  };

  const handleToTokenChange = (value: string): void => {
    if (value === fromToken) {
      setFromToken(toToken);
    }
    setToToken(value);
  };

  const handleAmountChange = (e: ChangeEvent<HTMLInputElement>): void => {
    const { value } = e.target;
    setAmount(value === "" ? null : parseFloat(value));
  };

  const handleConvertRates = (): void => {
    if (!fromToken || !toToken) {
      toast.error("Please select both currencies.");
      return;
    }

    if (!amount || amount <= 0) {
      toast.error("Please enter a valid amount.");
      return;
    }
    if (!prices[fromToken] || !prices[toToken]) {
      toast.error("Conversion rates not available.");
      return;
    }

    const rate = prices[fromToken] / prices[toToken];
    const result = amount * rate;
    setConvertedAmount(result.toFixed(2));

    toast.success(
      `Conversion successful: ${amount} ${fromToken} = ${result.toFixed(
        2
      )} ${toToken}`
    );

    setShowWalletOption(true);
  };

  const connectWallet = async (): Promise<boolean> => {
    return new Promise((resolve) => {
      setTimeout(() => {
        const isSuccess = Math.random() < 0.9;
        resolve(isSuccess);
      }, 1000);
    });
  };

  const handleConnectWallet = async (): Promise<void> => {
    toast.info("Connecting to e-wallet...");
    const walletConnected = await connectWallet();
    if (walletConnected) {
      toast.success("Wallet connected successfully!");
      setShowWalletOption(false);
    } else {
      toast.error("E-wallet connection failed. Please try again.");
    }
  };

  const exchangeRate =
    fromToken && toToken && prices ? prices[fromToken] / prices[toToken] : null;

  return (
    <>
      <Card className="w-fit h-full sm:p-2 md:p-4 space-y-2 back">
        <CardHeader className="text-center">
          <CardTitle className="text-2xl font-bold">Currency Swap</CardTitle>
          <CardDescription>Convert between multiple currencies</CardDescription>
        </CardHeader>
        <CardContent className="w-full p-2">
          {isLoading ? (
            <div className="flex justify-center">
              <ClipLoader className="w-6 h-6 text-blue-500" />
            </div>
          ) : error ? (
            <div className="text-red-500 text-center">{error}</div>
          ) : (
            <div>
              <div className="w-full">
                <label className="block text-sm font-medium mb-1">From</label>
                <div className="grid grid-cols-2 items-center gap-2">
                  <Input
                    type="number"
                    placeholder="Amount"
                    value={amount === null ? "" : amount}
                    onChange={handleAmountChange}
                    className="w-full"
                    id="from"
                  />
                  <Popover open={openFrom} onOpenChange={setOpenFrom}>
                    <PopoverTrigger asChild>
                      <Button
                        variant="outline"
                        role="combobox"
                        aria-expanded={openFrom}
                        className="w-full justify-between"
                      >
                        {fromToken ? (
                          <div className="flex items-center">
                            <Image
                              width={20}
                              height={20}
                              src={`/tokens/${fromToken}.svg`}
                              alt={fromToken}
                              className="inline mr-2"
                            />
                            <span>{fromToken}</span>
                          </div>
                        ) : (
                          "Select currency"
                        )}
                        <ChevronsUpDown className="opacity-50" />
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                      <Command>
                        <CommandInput
                          placeholder="Search currency"
                          value={fromSearchQuery}
                          onValueChange={setFromSearchQuery}
                        />
                        <CommandList>
                          <CommandEmpty>No currency found.</CommandEmpty>
                          <CommandGroup>
                            {tokens
                              .filter((token) =>
                                token.currency
                                  .toLowerCase()
                                  .includes(fromSearchQuery.toLowerCase())
                              )
                              .map((token) => (
                                <CommandItem
                                  key={token.currency}
                                  value={token.currency}
                                  onSelect={(currentValue) => {
                                    handleFromTokenChange(currentValue);
                                    setOpenFrom(false);
                                    setFromSearchQuery("");
                                  }}
                                >
                                  <Image
                                    width={20}
                                    height={20}
                                    src={`/tokens/${token.currency}.svg`}
                                    alt={token.currency}
                                    className="inline mr-2"
                                  />
                                  {token.currency}
                                  {fromToken === token.currency && (
                                    <Check className="ml-auto opacity-100" />
                                  )}
                                </CommandItem>
                              ))}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
                </div>
              </div>

              <div className="w-full my-2 flex items-center justify-center">
                <motion.button
                  whileHover={{ scale: 1.1 }}
                  whileTap={{ scale: 0.9 }}
                  className="p-2 rounded-full shadow-lg bg-gray-100"
                  onClick={handleSwap}
                >
                  ⬆⬇
                </motion.button>
              </div>

              {/* To Currency with Combobox */}
              <div className="w-full mb-4">
                <label className="block text-sm font-medium mb-1">To</label>
                <div className="grid grid-cols-2 items-center gap-2">
                  <div className="text-2xl font-bold">{convertedAmount}</div>
                  <Popover open={openTo} onOpenChange={setOpenTo}>
                    <PopoverTrigger asChild>
                      <Button
                        variant="outline"
                        role="combobox"
                        aria-expanded={openTo}
                        className="w-full justify-between"
                      >
                        {toToken ? (
                          <div className="flex items-center">
                            <Image
                              width={20}
                              height={20}
                              src={`/tokens/${toToken}.svg`}
                              alt={toToken}
                              className="inline mr-2"
                            />
                            <span>{toToken}</span>
                          </div>
                        ) : (
                          "Select currency"
                        )}
                        <ChevronsUpDown className="opacity-50" />
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                      <Command>
                        <CommandInput
                          placeholder="Search currency..."
                          value={toSearchQuery}
                          onValueChange={setToSearchQuery}
                        />
                        <CommandList>
                          <CommandEmpty>No currency found.</CommandEmpty>
                          <CommandGroup>
                            {tokens
                              .filter((token) =>
                                token.currency
                                  .toLowerCase()
                                  .includes(toSearchQuery.toLowerCase())
                              )
                              .map((token) => (
                                <CommandItem
                                  key={token.currency}
                                  value={token.currency}
                                  onSelect={(currentValue) => {
                                    handleToTokenChange(currentValue);
                                    setOpenTo(false);
                                    setToSearchQuery("");
                                  }}
                                >
                                  <Image
                                    width={20}
                                    height={20}
                                    src={`/tokens/${token.currency}.svg`}
                                    alt={token.currency}
                                    className="inline mr-2"
                                  />
                                  {token.currency}
                                  {toToken === token.currency && (
                                    <Check className="ml-auto opacity-100" />
                                  )}
                                </CommandItem>
                              ))}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
                </div>
              </div>
            </div>
          )}
        </CardContent>
        <CardFooter>
          <div className="flex w-full flex-col">
            {fromToken && toToken && exchangeRate && (
              <div className="text-sm text-gray-300">
                1 <span className="text-blue-400">{fromToken}</span> ={" "}
                {exchangeRate.toFixed(8)}{" "}
                <span className="text-green-400">{toToken}</span>
              </div>
            )}
            <Button
              className="bg-green-200 backdrop-blur-sm w-full p-3 rounded-lg text-white font-bold hover:bg-green-600"
              onClick={handleConvertRates}
            >
              <span className="font-bold text-green-400">Convert Now</span>
            </Button>

            {showWalletOption && (
              <div className="mt-4">
                <p className="text-sm text-gray-300">
                  Conversion complete. Would you like to connect to your
                  e-wallet?
                </p>
                <Button
                  className="bg-blue-500 w-full p-3 rounded-lg text-white font-bold hover:bg-blue-600 mt-2"
                  onClick={handleConnectWallet}
                >
                  Connect Wallet
                </Button>
              </div>
            )}
          </div>
        </CardFooter>
      </Card>
      <ToastContainer position="top-right" autoClose={3000} />
    </>
  );
}
