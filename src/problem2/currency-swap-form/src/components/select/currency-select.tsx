import { Dispatch, SetStateAction } from "react";
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
import Image from "next/image";
import { Button } from "../ui/button";
import { Check, ChevronsUpDown } from "lucide-react";
import { Token } from "@/types/types";

type CurrencySelectProps = Readonly<{
  open: boolean;
  setOpen: Dispatch<SetStateAction<boolean>>;
  currentToken: string;
  tokens: Token[];
  handleTokenChange: (value: string) => void;
  setSearchQuery: Dispatch<SetStateAction<string>>;
  searchQuery: string;
}>;

export default function CurrencySelect({
  open,
  setOpen,
  currentToken,
  tokens,
  handleTokenChange,
  searchQuery,
  setSearchQuery,
}: CurrencySelectProps) {
  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {currentToken ? (
            <div className="flex items-center">
              <Image
                width={20}
                height={20}
                src={`/tokens/${currentToken}.svg`}
                alt={currentToken}
                className="inline mr-2"
              />
              <span>{currentToken}</span>
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
            value={searchQuery}
            onValueChange={setSearchQuery}
          />
          <CommandList>
            <CommandEmpty>No currency found.</CommandEmpty>
            <CommandGroup>
              {tokens
                .filter((token) =>
                  token.currency
                    .toLowerCase()
                    .includes(searchQuery.toLowerCase())
                )
                .map((token) => (
                  <CommandItem
                    key={token.currency}
                    value={token.currency}
                    onSelect={(currentValue) => {
                      handleTokenChange(currentValue);
                      setOpen(false);
                      setSearchQuery("");
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
                    {currentToken === token.currency && (
                      <Check className="ml-auto opacity-100" />
                    )}
                  </CommandItem>
                ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
