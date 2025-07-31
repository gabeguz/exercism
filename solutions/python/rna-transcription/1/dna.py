def to_rna(dna_strand):
    rna_strand = []
    for dna_nucleotide in dna_strand:
        if dna_nucleotide == 'G':
            rna_nucleotide = 'C'
        if dna_nucleotide == 'C':
            rna_nucleotide = 'G'
        if dna_nucleotide == 'A':
            rna_nucleotide = 'U'
        if dna_nucleotide == 'T':
            rna_nucleotide = 'A'
        rna_strand.append(rna_nucleotide)

    return ''.join(rna_strand)